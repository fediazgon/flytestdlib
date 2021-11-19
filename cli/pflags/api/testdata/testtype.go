// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package api

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (TestType) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (TestType) mustJsonMarshal(v interface{}) string {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(raw)
}

func (TestType) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in TestType and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg TestType) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("TestType", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "str"), DefaultTestType.StringValue, "life is short")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "bl"), DefaultTestType.BoolValue, "")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "nested.i"), DefaultTestType.NestedType.IntValue, "this is an important flag")
	cmdFlags.IntSlice(fmt.Sprintf("%v%v", prefix, "ints"), []int{12, 1}, "")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "strs"), []string{"12", "1"}, "")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "complexArr"), []string{}, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "c"), DefaultTestType.mustJsonMarshal(DefaultTestType.StringToJSON), "I'm a complex type but can be converted from string.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage.type"), DefaultTestType.StorageConfig.Type, "Sets the type of storage to configure [s3/minio/local/mem/stow].")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage.connection.endpoint"), DefaultTestType.StorageConfig.Connection.Endpoint.String(), "URL for storage client to connect to.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage.connection.auth-type"), DefaultTestType.StorageConfig.Connection.AuthType, "Auth Type to use [iam, accesskey].")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage.connection.access-key"), DefaultTestType.StorageConfig.Connection.AccessKey, "Access key to use. Only required when authtype is set to accesskey.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage.connection.secret-key"), DefaultTestType.StorageConfig.Connection.SecretKey, "Secret to use when accesskey is set.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage.connection.region"), DefaultTestType.StorageConfig.Connection.Region, "Region to connect to.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "storage.connection.disable-ssl"), DefaultTestType.StorageConfig.Connection.DisableSSL, "Disables SSL connection. Should only be used for development.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage.stow.kind"), DefaultTestType.StorageConfig.Stow.Kind, "Kind of Stow backend to use. Refer to github/graymeta/stow")
	cmdFlags.StringToString(fmt.Sprintf("%v%v", prefix, "storage.stow.config"), DefaultTestType.StorageConfig.Stow.Config, "Configuration for stow backend. Refer to github/graymeta/stow")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage.container"), DefaultTestType.StorageConfig.InitContainer, "Initial container (in s3 a bucket) to create -if it doesn't exist-.'")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "storage.enable-multicontainer"), DefaultTestType.StorageConfig.MultiContainerEnabled, "If this is true,  then the container argument is overlooked and redundant. This config will automatically open new connections to new containers/buckets as they are encountered")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "storage.cache.max_size_mbs"), DefaultTestType.StorageConfig.Cache.MaxSizeMegabytes, "Maximum size of the cache where the Blob store data is cached in-memory. If not specified or set to 0,  cache is not used")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "storage.cache.target_gc_percent"), DefaultTestType.StorageConfig.Cache.TargetGCPercent, "Sets the garbage collection target percentage.")
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "storage.limits.maxDownloadMBs"), DefaultTestType.StorageConfig.Limits.GetLimitMegabytes, "Maximum allowed download size (in MBs) per call.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage.defaultHttpClient.timeout"), DefaultTestType.StorageConfig.DefaultHTTPClient.Timeout.String(), "Sets time out on the http client.")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "i"), DefaultTestType.elemValueOrNil(DefaultTestType.IntValue).(int), "")
	cmdFlags.StringToString(fmt.Sprintf("%v%v", prefix, "m"), DefaultTestType.StringMap, "I'm a map of strings")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "constType"), int(DefaultTestType.ConstType), "")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "aliasType"), DefaultTestType.AliasType, "")
	return cmdFlags
}
