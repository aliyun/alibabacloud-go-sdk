// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CreateUserDeliveryTaskRequest
	GetBusinessType() *string
	SetDataCenter(v string) *CreateUserDeliveryTaskRequest
	GetDataCenter() *string
	SetDeliveryType(v string) *CreateUserDeliveryTaskRequest
	GetDeliveryType() *string
	SetDetails(v string) *CreateUserDeliveryTaskRequest
	GetDetails() *string
	SetDiscardRate(v float32) *CreateUserDeliveryTaskRequest
	GetDiscardRate() *float32
	SetFieldName(v string) *CreateUserDeliveryTaskRequest
	GetFieldName() *string
	SetFilterVer(v string) *CreateUserDeliveryTaskRequest
	GetFilterVer() *string
	SetHttpDelivery(v *CreateUserDeliveryTaskRequestHttpDelivery) *CreateUserDeliveryTaskRequest
	GetHttpDelivery() *CreateUserDeliveryTaskRequestHttpDelivery
	SetKafkaDelivery(v *CreateUserDeliveryTaskRequestKafkaDelivery) *CreateUserDeliveryTaskRequest
	GetKafkaDelivery() *CreateUserDeliveryTaskRequestKafkaDelivery
	SetOssDelivery(v *CreateUserDeliveryTaskRequestOssDelivery) *CreateUserDeliveryTaskRequest
	GetOssDelivery() *CreateUserDeliveryTaskRequestOssDelivery
	SetS3Delivery(v *CreateUserDeliveryTaskRequestS3Delivery) *CreateUserDeliveryTaskRequest
	GetS3Delivery() *CreateUserDeliveryTaskRequestS3Delivery
	SetSlsDelivery(v *CreateUserDeliveryTaskRequestSlsDelivery) *CreateUserDeliveryTaskRequest
	GetSlsDelivery() *CreateUserDeliveryTaskRequestSlsDelivery
	SetTaskName(v string) *CreateUserDeliveryTaskRequest
	GetTaskName() *string
}

type CreateUserDeliveryTaskRequest struct {
	// The log category. Valid values:
	//
	// 	- dcdn_log_access_l1 (default): access logs.
	//
	// 	- dcdn_log_er: Edge Routine logs.
	//
	// 	- dcdn_log_waf: firewall logs.
	//
	// 	- dcdn_log_ipa: TCP/UDP proxy logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The data center. Valid values:
	//
	// 	- cn: the Chinese mainland.
	//
	// 	- sg: outside the Chinese mainland.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The destination of the delivery. Valid values:
	//
	// 1.  sls: Alibaba Cloud SLS.
	//
	// 2.  http: HTTP server.
	//
	// 3.  aws3: Amazon S3.
	//
	// 4.  oss: Alibaba Cloud OSS.
	//
	// 5.  kafka: Kafka.
	//
	// 6.  aws3cmpt: S3-compatible storage service.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	Details      *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// The discard rate. Default value: 0.
	//
	// example:
	//
	// 0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// The log field. If you specify multiple fields, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// user_agent,ip_address,ip_port
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FilterVer *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	// The configurations for delivery to an HTTP server.
	HttpDelivery *CreateUserDeliveryTaskRequestHttpDelivery `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty" type:"Struct"`
	// The configurations for delivery to Kafka.
	KafkaDelivery *CreateUserDeliveryTaskRequestKafkaDelivery `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty" type:"Struct"`
	// The configurations for delivery to OSS.
	OssDelivery *CreateUserDeliveryTaskRequestOssDelivery `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty" type:"Struct"`
	// The configurations for delivery to Amazon S3 or an S3-compatible service.
	S3Delivery *CreateUserDeliveryTaskRequestS3Delivery `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty" type:"Struct"`
	// The configurations for delivery to SLS.
	SlsDelivery *CreateUserDeliveryTaskRequestSlsDelivery `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty" type:"Struct"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateUserDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateUserDeliveryTaskRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateUserDeliveryTaskRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *CreateUserDeliveryTaskRequest) GetDetails() *string {
	return s.Details
}

func (s *CreateUserDeliveryTaskRequest) GetDiscardRate() *float32 {
	return s.DiscardRate
}

func (s *CreateUserDeliveryTaskRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateUserDeliveryTaskRequest) GetFilterVer() *string {
	return s.FilterVer
}

func (s *CreateUserDeliveryTaskRequest) GetHttpDelivery() *CreateUserDeliveryTaskRequestHttpDelivery {
	return s.HttpDelivery
}

func (s *CreateUserDeliveryTaskRequest) GetKafkaDelivery() *CreateUserDeliveryTaskRequestKafkaDelivery {
	return s.KafkaDelivery
}

func (s *CreateUserDeliveryTaskRequest) GetOssDelivery() *CreateUserDeliveryTaskRequestOssDelivery {
	return s.OssDelivery
}

func (s *CreateUserDeliveryTaskRequest) GetS3Delivery() *CreateUserDeliveryTaskRequestS3Delivery {
	return s.S3Delivery
}

func (s *CreateUserDeliveryTaskRequest) GetSlsDelivery() *CreateUserDeliveryTaskRequestSlsDelivery {
	return s.SlsDelivery
}

func (s *CreateUserDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateUserDeliveryTaskRequest) SetBusinessType(v string) *CreateUserDeliveryTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetDataCenter(v string) *CreateUserDeliveryTaskRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetDeliveryType(v string) *CreateUserDeliveryTaskRequest {
	s.DeliveryType = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetDetails(v string) *CreateUserDeliveryTaskRequest {
	s.Details = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetDiscardRate(v float32) *CreateUserDeliveryTaskRequest {
	s.DiscardRate = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetFieldName(v string) *CreateUserDeliveryTaskRequest {
	s.FieldName = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetFilterVer(v string) *CreateUserDeliveryTaskRequest {
	s.FilterVer = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetHttpDelivery(v *CreateUserDeliveryTaskRequestHttpDelivery) *CreateUserDeliveryTaskRequest {
	s.HttpDelivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetKafkaDelivery(v *CreateUserDeliveryTaskRequestKafkaDelivery) *CreateUserDeliveryTaskRequest {
	s.KafkaDelivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetOssDelivery(v *CreateUserDeliveryTaskRequestOssDelivery) *CreateUserDeliveryTaskRequest {
	s.OssDelivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetS3Delivery(v *CreateUserDeliveryTaskRequestS3Delivery) *CreateUserDeliveryTaskRequest {
	s.S3Delivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetSlsDelivery(v *CreateUserDeliveryTaskRequestSlsDelivery) *CreateUserDeliveryTaskRequest {
	s.SlsDelivery = v
	return s
}

func (s *CreateUserDeliveryTaskRequest) SetTaskName(v string) *CreateUserDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateUserDeliveryTaskRequest) Validate() error {
	if s.HttpDelivery != nil {
		if err := s.HttpDelivery.Validate(); err != nil {
			return err
		}
	}
	if s.KafkaDelivery != nil {
		if err := s.KafkaDelivery.Validate(); err != nil {
			return err
		}
	}
	if s.OssDelivery != nil {
		if err := s.OssDelivery.Validate(); err != nil {
			return err
		}
	}
	if s.S3Delivery != nil {
		if err := s.S3Delivery.Validate(); err != nil {
			return err
		}
	}
	if s.SlsDelivery != nil {
		if err := s.SlsDelivery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserDeliveryTaskRequestHttpDelivery struct {
	// The compression method.
	//
	// example:
	//
	// gzip
	Compress *string `json:"Compress,omitempty" xml:"Compress,omitempty"`
	// The address of the HTTP server.
	//
	// example:
	//
	// http://xxx.aliyun.com/v1/log/upload
	DestUrl *string `json:"DestUrl,omitempty" xml:"DestUrl,omitempty"`
	// The custom headers.
	HeaderParam map[string]*HttpDeliveryHeaderParamValue `json:"HeaderParam,omitempty" xml:"HeaderParam,omitempty"`
	// The ending separator.
	//
	// example:
	//
	// \\n
	LastLogSplit *bool `json:"LastLogSplit,omitempty" xml:"LastLogSplit,omitempty"`
	// The prefix of the log delivery package.
	//
	// example:
	//
	// cdnVersion:1.0
	LogBodyPrefix *string `json:"LogBodyPrefix,omitempty" xml:"LogBodyPrefix,omitempty"`
	// The suffix of the log delivery package.
	//
	// example:
	//
	// cdnVersion:1.0
	LogBodySuffix *string `json:"LogBodySuffix,omitempty" xml:"LogBodySuffix,omitempty"`
	// Specifies whether to enable log splitting. Default value: true.
	//
	// example:
	//
	// true
	LogSplit *bool `json:"LogSplit,omitempty" xml:"LogSplit,omitempty"`
	// The log separator.
	//
	// example:
	//
	// \\n
	LogSplitWords *string `json:"LogSplitWords,omitempty" xml:"LogSplitWords,omitempty"`
	// The maximum size of data for each delivery. Unit: MB.
	//
	// example:
	//
	// 5
	MaxBatchMB *int64 `json:"MaxBatchMB,omitempty" xml:"MaxBatchMB,omitempty"`
	// The maximum number of entries for each delivery.
	//
	// example:
	//
	// 1000
	MaxBatchSize *int64 `json:"MaxBatchSize,omitempty" xml:"MaxBatchSize,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 3
	MaxRetry *int64 `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	// The custom query parameters.
	QueryParam map[string]*HttpDeliveryQueryParamValue `json:"QueryParam,omitempty" xml:"QueryParam,omitempty"`
	// Specifies whether to use server authentication.
	//
	// example:
	//
	// true
	StandardAuthOn *bool `json:"StandardAuthOn,omitempty" xml:"StandardAuthOn,omitempty"`
	// The authentication configurations.
	StandardAuthParam *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam `json:"StandardAuthParam,omitempty" xml:"StandardAuthParam,omitempty" type:"Struct"`
	// The timeout period. Unit: seconds.
	//
	// example:
	//
	// 10
	TransformTimeout *int64 `json:"TransformTimeout,omitempty" xml:"TransformTimeout,omitempty"`
}

func (s CreateUserDeliveryTaskRequestHttpDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestHttpDelivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetCompress() *string {
	return s.Compress
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetDestUrl() *string {
	return s.DestUrl
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetHeaderParam() map[string]*HttpDeliveryHeaderParamValue {
	return s.HeaderParam
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetLastLogSplit() *bool {
	return s.LastLogSplit
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetLogBodyPrefix() *string {
	return s.LogBodyPrefix
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetLogBodySuffix() *string {
	return s.LogBodySuffix
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetLogSplit() *bool {
	return s.LogSplit
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetLogSplitWords() *string {
	return s.LogSplitWords
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetMaxBatchMB() *int64 {
	return s.MaxBatchMB
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetMaxBatchSize() *int64 {
	return s.MaxBatchSize
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetMaxRetry() *int64 {
	return s.MaxRetry
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetQueryParam() map[string]*HttpDeliveryQueryParamValue {
	return s.QueryParam
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetStandardAuthOn() *bool {
	return s.StandardAuthOn
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetStandardAuthParam() *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	return s.StandardAuthParam
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) GetTransformTimeout() *int64 {
	return s.TransformTimeout
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetCompress(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.Compress = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetDestUrl(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.DestUrl = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetHeaderParam(v map[string]*HttpDeliveryHeaderParamValue) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.HeaderParam = v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLastLogSplit(v bool) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LastLogSplit = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLogBodyPrefix(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LogBodyPrefix = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLogBodySuffix(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LogBodySuffix = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLogSplit(v bool) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LogSplit = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetLogSplitWords(v string) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.LogSplitWords = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetMaxBatchMB(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.MaxBatchMB = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetMaxBatchSize(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.MaxBatchSize = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetMaxRetry(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.MaxRetry = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetQueryParam(v map[string]*HttpDeliveryQueryParamValue) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.QueryParam = v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetStandardAuthOn(v bool) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.StandardAuthOn = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetStandardAuthParam(v *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.StandardAuthParam = v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) SetTransformTimeout(v int64) *CreateUserDeliveryTaskRequestHttpDelivery {
	s.TransformTimeout = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDelivery) Validate() error {
	if s.StandardAuthParam != nil {
		if err := s.StandardAuthParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam struct {
	// The validity period of the signature.
	//
	// example:
	//
	// 300
	ExpiredTime *int32 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The private key.
	//
	// example:
	//
	// xxxx
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The URL path.
	//
	// example:
	//
	// v1/log/upload
	UrlPath *string `json:"UrlPath,omitempty" xml:"UrlPath,omitempty"`
}

func (s CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) GetExpiredTime() *int32 {
	return s.ExpiredTime
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) GetUrlPath() *string {
	return s.UrlPath
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetExpiredTime(v int32) *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.ExpiredTime = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetPrivateKey(v string) *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.PrivateKey = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetUrlPath(v string) *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.UrlPath = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestHttpDeliveryStandardAuthParam) Validate() error {
	return dara.Validate(s)
}

type CreateUserDeliveryTaskRequestKafkaDelivery struct {
	// The load balancing method.
	//
	// example:
	//
	// kafka.LeastBytes
	Balancer *string `json:"Balancer,omitempty" xml:"Balancer,omitempty"`
	// The brokers.
	Brokers []*string `json:"Brokers,omitempty" xml:"Brokers,omitempty" type:"Repeated"`
	// The compression method. By default, data is not compressed.
	//
	// example:
	//
	// gzip
	Compress *string `json:"Compress,omitempty" xml:"Compress,omitempty"`
	// The encryption method.
	//
	// example:
	//
	// plain
	MachanismType *string `json:"MachanismType,omitempty" xml:"MachanismType,omitempty"`
	// The password.
	//
	// example:
	//
	// xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The topic.
	//
	// example:
	//
	// dqc_test2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// Specifies whether to enable authentication.
	//
	// example:
	//
	// true
	UserAuth *bool `json:"UserAuth,omitempty" xml:"UserAuth,omitempty"`
	// The username.
	//
	// example:
	//
	// xxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserDeliveryTaskRequestKafkaDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestKafkaDelivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) GetBalancer() *string {
	return s.Balancer
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) GetBrokers() []*string {
	return s.Brokers
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) GetCompress() *string {
	return s.Compress
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) GetMachanismType() *string {
	return s.MachanismType
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) GetPassword() *string {
	return s.Password
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) GetTopic() *string {
	return s.Topic
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) GetUserAuth() *bool {
	return s.UserAuth
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetBalancer(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Balancer = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetBrokers(v []*string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Brokers = v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetCompress(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Compress = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetMachanismType(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.MachanismType = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetPassword(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Password = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetTopic(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.Topic = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetUserAuth(v bool) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.UserAuth = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) SetUserName(v string) *CreateUserDeliveryTaskRequestKafkaDelivery {
	s.UserName = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestKafkaDelivery) Validate() error {
	return dara.Validate(s)
}

type CreateUserDeliveryTaskRequestOssDelivery struct {
	// The ID of your Alibaba Cloud account.
	//
	// example:
	//
	// 1234***
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// test_rlog
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The prefix of the path in which you want to store logs.
	//
	// example:
	//
	// logriver-test/log
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// The region in which the bucket is located.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateUserDeliveryTaskRequestOssDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestOssDelivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) GetAliuid() *string {
	return s.Aliuid
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) GetPrefixPath() *string {
	return s.PrefixPath
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) GetRegion() *string {
	return s.Region
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) SetAliuid(v string) *CreateUserDeliveryTaskRequestOssDelivery {
	s.Aliuid = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) SetBucketName(v string) *CreateUserDeliveryTaskRequestOssDelivery {
	s.BucketName = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) SetPrefixPath(v string) *CreateUserDeliveryTaskRequestOssDelivery {
	s.PrefixPath = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) SetRegion(v string) *CreateUserDeliveryTaskRequestOssDelivery {
	s.Region = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestOssDelivery) Validate() error {
	return dara.Validate(s)
}

type CreateUserDeliveryTaskRequestS3Delivery struct {
	// The access key ID of your Amazon S3 account.
	//
	// example:
	//
	// g0f46623ll0g0
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The directory in the bucket.
	//
	// example:
	//
	// logriver-test/log
	BucketPath *string `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// https://s3.oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The prefix of the path in which you want to store logs.
	//
	// example:
	//
	// logriver-test/log
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// The region ID of the service.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Specifies whether the service is compatible with Amazon S3.
	//
	// example:
	//
	// true
	S3Cmpt *bool `json:"S3Cmpt,omitempty" xml:"S3Cmpt,omitempty"`
	// The secret access key of your Amazon S3 account.
	//
	// example:
	//
	// ***
	SecretKey            *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	ServerSideEncryption *bool   `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	VertifyType          *string `json:"VertifyType,omitempty" xml:"VertifyType,omitempty"`
}

func (s CreateUserDeliveryTaskRequestS3Delivery) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestS3Delivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) GetBucketPath() *string {
	return s.BucketPath
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) GetPrefixPath() *string {
	return s.PrefixPath
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) GetRegion() *string {
	return s.Region
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) GetS3Cmpt() *bool {
	return s.S3Cmpt
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) GetSecretKey() *string {
	return s.SecretKey
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) GetServerSideEncryption() *bool {
	return s.ServerSideEncryption
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) GetVertifyType() *string {
	return s.VertifyType
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetAccessKey(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.AccessKey = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetBucketPath(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.BucketPath = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetEndpoint(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.Endpoint = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetPrefixPath(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.PrefixPath = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetRegion(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.Region = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetS3Cmpt(v bool) *CreateUserDeliveryTaskRequestS3Delivery {
	s.S3Cmpt = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetSecretKey(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.SecretKey = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetServerSideEncryption(v bool) *CreateUserDeliveryTaskRequestS3Delivery {
	s.ServerSideEncryption = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) SetVertifyType(v string) *CreateUserDeliveryTaskRequestS3Delivery {
	s.VertifyType = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestS3Delivery) Validate() error {
	return dara.Validate(s)
}

type CreateUserDeliveryTaskRequestSlsDelivery struct {
	// The name of the SLS Logstore.
	//
	// example:
	//
	// accesslog-test
	SLSLogStore *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	// The name of the SLS project.
	//
	// example:
	//
	// dcdn-test20240417
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	// The region in which the SLS project resides.
	//
	// example:
	//
	// cn-hangzhou
	SLSRegion *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
}

func (s CreateUserDeliveryTaskRequestSlsDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskRequestSlsDelivery) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) GetSLSLogStore() *string {
	return s.SLSLogStore
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) GetSLSProject() *string {
	return s.SLSProject
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) GetSLSRegion() *string {
	return s.SLSRegion
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) SetSLSLogStore(v string) *CreateUserDeliveryTaskRequestSlsDelivery {
	s.SLSLogStore = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) SetSLSProject(v string) *CreateUserDeliveryTaskRequestSlsDelivery {
	s.SLSProject = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) SetSLSRegion(v string) *CreateUserDeliveryTaskRequestSlsDelivery {
	s.SLSRegion = &v
	return s
}

func (s *CreateUserDeliveryTaskRequestSlsDelivery) Validate() error {
	return dara.Validate(s)
}
