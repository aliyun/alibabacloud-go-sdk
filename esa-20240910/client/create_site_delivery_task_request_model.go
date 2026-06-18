// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CreateSiteDeliveryTaskRequest
	GetBusinessType() *string
	SetDataCenter(v string) *CreateSiteDeliveryTaskRequest
	GetDataCenter() *string
	SetDeliveryType(v string) *CreateSiteDeliveryTaskRequest
	GetDeliveryType() *string
	SetDiscardRate(v float32) *CreateSiteDeliveryTaskRequest
	GetDiscardRate() *float32
	SetFieldName(v string) *CreateSiteDeliveryTaskRequest
	GetFieldName() *string
	SetFilterVer(v string) *CreateSiteDeliveryTaskRequest
	GetFilterVer() *string
	SetHttpDelivery(v *CreateSiteDeliveryTaskRequestHttpDelivery) *CreateSiteDeliveryTaskRequest
	GetHttpDelivery() *CreateSiteDeliveryTaskRequestHttpDelivery
	SetKafkaDelivery(v *CreateSiteDeliveryTaskRequestKafkaDelivery) *CreateSiteDeliveryTaskRequest
	GetKafkaDelivery() *CreateSiteDeliveryTaskRequestKafkaDelivery
	SetOssDelivery(v *CreateSiteDeliveryTaskRequestOssDelivery) *CreateSiteDeliveryTaskRequest
	GetOssDelivery() *CreateSiteDeliveryTaskRequestOssDelivery
	SetS3Delivery(v *CreateSiteDeliveryTaskRequestS3Delivery) *CreateSiteDeliveryTaskRequest
	GetS3Delivery() *CreateSiteDeliveryTaskRequestS3Delivery
	SetSiteId(v int64) *CreateSiteDeliveryTaskRequest
	GetSiteId() *int64
	SetSlsDelivery(v *CreateSiteDeliveryTaskRequestSlsDelivery) *CreateSiteDeliveryTaskRequest
	GetSlsDelivery() *CreateSiteDeliveryTaskRequestSlsDelivery
	SetTaskName(v string) *CreateSiteDeliveryTaskRequest
	GetTaskName() *string
}

type CreateSiteDeliveryTaskRequest struct {
	// The business type. Valid values:
	//
	// - **dcdn_log_access_l1*	- (default): access log.
	//
	// - **dcdn_log_er**: edge function log.
	//
	// - **dcdn_log_waf**: WAF protection log.
	//
	// - **dcdn_log_ipa**: Layer-4 acceleration log.
	//
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The data center. Valid values:
	//
	// - **cn**: Chinese mainland.
	//
	// - **oversea**: regions outside the Chinese mainland.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The type of the delivery destination. Valid values:
	//
	// - **sls**: Log Service.
	//
	// - **http**: an HTTP service.
	//
	// - **aws3**: Amazon S3.
	//
	// - **oss**: Object Storage Service.
	//
	// - **kafka**: a Kafka service.
	//
	// - **aws3cmpt**: an Amazon S3-compatible service.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The discard rate. If you do not specify this parameter, the default value 0 is used.
	//
	// example:
	//
	// 0.0
	DiscardRate *float32 `json:"DiscardRate,omitempty" xml:"DiscardRate,omitempty"`
	// The log fields for delivery. Separate multiple fields with a comma (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// user_agent,ip_adress,ip_port
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FilterVer *string `json:"FilterVer,omitempty" xml:"FilterVer,omitempty"`
	// The parameters for delivering logs to an HTTP server.
	HttpDelivery *CreateSiteDeliveryTaskRequestHttpDelivery `json:"HttpDelivery,omitempty" xml:"HttpDelivery,omitempty" type:"Struct"`
	// The parameters for delivering logs to a Kafka cluster.
	KafkaDelivery *CreateSiteDeliveryTaskRequestKafkaDelivery `json:"KafkaDelivery,omitempty" xml:"KafkaDelivery,omitempty" type:"Struct"`
	// The parameters for delivering logs to Object Storage Service (OSS).
	OssDelivery *CreateSiteDeliveryTaskRequestOssDelivery `json:"OssDelivery,omitempty" xml:"OssDelivery,omitempty" type:"Struct"`
	// The parameters for delivering logs to an Amazon S3 bucket or an S3-compatible service.
	S3Delivery *CreateSiteDeliveryTaskRequestS3Delivery `json:"S3Delivery,omitempty" xml:"S3Delivery,omitempty" type:"Struct"`
	// The ID of the site. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12312312112***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The parameters for delivering logs to Log Service.
	SlsDelivery *CreateSiteDeliveryTaskRequestSlsDelivery `json:"SlsDelivery,omitempty" xml:"SlsDelivery,omitempty" type:"Struct"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dcdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateSiteDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateSiteDeliveryTaskRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateSiteDeliveryTaskRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *CreateSiteDeliveryTaskRequest) GetDiscardRate() *float32 {
	return s.DiscardRate
}

func (s *CreateSiteDeliveryTaskRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateSiteDeliveryTaskRequest) GetFilterVer() *string {
	return s.FilterVer
}

func (s *CreateSiteDeliveryTaskRequest) GetHttpDelivery() *CreateSiteDeliveryTaskRequestHttpDelivery {
	return s.HttpDelivery
}

func (s *CreateSiteDeliveryTaskRequest) GetKafkaDelivery() *CreateSiteDeliveryTaskRequestKafkaDelivery {
	return s.KafkaDelivery
}

func (s *CreateSiteDeliveryTaskRequest) GetOssDelivery() *CreateSiteDeliveryTaskRequestOssDelivery {
	return s.OssDelivery
}

func (s *CreateSiteDeliveryTaskRequest) GetS3Delivery() *CreateSiteDeliveryTaskRequestS3Delivery {
	return s.S3Delivery
}

func (s *CreateSiteDeliveryTaskRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateSiteDeliveryTaskRequest) GetSlsDelivery() *CreateSiteDeliveryTaskRequestSlsDelivery {
	return s.SlsDelivery
}

func (s *CreateSiteDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateSiteDeliveryTaskRequest) SetBusinessType(v string) *CreateSiteDeliveryTaskRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetDataCenter(v string) *CreateSiteDeliveryTaskRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetDeliveryType(v string) *CreateSiteDeliveryTaskRequest {
	s.DeliveryType = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetDiscardRate(v float32) *CreateSiteDeliveryTaskRequest {
	s.DiscardRate = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetFieldName(v string) *CreateSiteDeliveryTaskRequest {
	s.FieldName = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetFilterVer(v string) *CreateSiteDeliveryTaskRequest {
	s.FilterVer = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetHttpDelivery(v *CreateSiteDeliveryTaskRequestHttpDelivery) *CreateSiteDeliveryTaskRequest {
	s.HttpDelivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetKafkaDelivery(v *CreateSiteDeliveryTaskRequestKafkaDelivery) *CreateSiteDeliveryTaskRequest {
	s.KafkaDelivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetOssDelivery(v *CreateSiteDeliveryTaskRequestOssDelivery) *CreateSiteDeliveryTaskRequest {
	s.OssDelivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetS3Delivery(v *CreateSiteDeliveryTaskRequestS3Delivery) *CreateSiteDeliveryTaskRequest {
	s.S3Delivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetSiteId(v int64) *CreateSiteDeliveryTaskRequest {
	s.SiteId = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetSlsDelivery(v *CreateSiteDeliveryTaskRequestSlsDelivery) *CreateSiteDeliveryTaskRequest {
	s.SlsDelivery = v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) SetTaskName(v string) *CreateSiteDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequest) Validate() error {
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

type CreateSiteDeliveryTaskRequestHttpDelivery struct {
	// The compression method. By default, logs are delivered uncompressed.
	//
	// example:
	//
	// gzip
	Compress *string `json:"Compress,omitempty" xml:"Compress,omitempty"`
	// The URL of the destination HTTP server.
	//
	// example:
	//
	// http://xxx.aliyun.com/v1/log/upload
	DestUrl *string `json:"DestUrl,omitempty" xml:"DestUrl,omitempty"`
	// The custom HTTP request headers.
	HeaderParam  map[string]*HttpDeliveryHeaderParamValue `json:"HeaderParam,omitempty" xml:"HeaderParam,omitempty"`
	LastLogSplit *bool                                    `json:"LastLogSplit,omitempty" xml:"LastLogSplit,omitempty"`
	// The prefix of the delivered log package.
	//
	// example:
	//
	// cdnVersion:1.0
	LogBodyPrefix *string `json:"LogBodyPrefix,omitempty" xml:"LogBodyPrefix,omitempty"`
	// The suffix of the delivered log package.
	//
	// example:
	//
	// cdnVersion:1.0
	LogBodySuffix *string `json:"LogBodySuffix,omitempty" xml:"LogBodySuffix,omitempty"`
	LogSplit      *bool   `json:"LogSplit,omitempty" xml:"LogSplit,omitempty"`
	LogSplitWords *string `json:"LogSplitWords,omitempty" xml:"LogSplitWords,omitempty"`
	// The maximum size of a delivery, in MB.
	//
	// example:
	//
	// 5
	MaxBatchMB *int64 `json:"MaxBatchMB,omitempty" xml:"MaxBatchMB,omitempty"`
	// The maximum number of log entries per delivery.
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
	// Specifies whether to use standard authentication.
	//
	// example:
	//
	// true
	StandardAuthOn *bool `json:"StandardAuthOn,omitempty" xml:"StandardAuthOn,omitempty"`
	// The standard authentication parameters.
	StandardAuthParam *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam `json:"StandardAuthParam,omitempty" xml:"StandardAuthParam,omitempty" type:"Struct"`
	// The timeout period, in seconds.
	//
	// example:
	//
	// 10
	TransformTimeout *int64 `json:"TransformTimeout,omitempty" xml:"TransformTimeout,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestHttpDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestHttpDelivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetCompress() *string {
	return s.Compress
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetDestUrl() *string {
	return s.DestUrl
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetHeaderParam() map[string]*HttpDeliveryHeaderParamValue {
	return s.HeaderParam
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetLastLogSplit() *bool {
	return s.LastLogSplit
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetLogBodyPrefix() *string {
	return s.LogBodyPrefix
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetLogBodySuffix() *string {
	return s.LogBodySuffix
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetLogSplit() *bool {
	return s.LogSplit
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetLogSplitWords() *string {
	return s.LogSplitWords
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetMaxBatchMB() *int64 {
	return s.MaxBatchMB
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetMaxBatchSize() *int64 {
	return s.MaxBatchSize
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetMaxRetry() *int64 {
	return s.MaxRetry
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetQueryParam() map[string]*HttpDeliveryQueryParamValue {
	return s.QueryParam
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetStandardAuthOn() *bool {
	return s.StandardAuthOn
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetStandardAuthParam() *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	return s.StandardAuthParam
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) GetTransformTimeout() *int64 {
	return s.TransformTimeout
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetCompress(v string) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.Compress = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetDestUrl(v string) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.DestUrl = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetHeaderParam(v map[string]*HttpDeliveryHeaderParamValue) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.HeaderParam = v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetLastLogSplit(v bool) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.LastLogSplit = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetLogBodyPrefix(v string) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.LogBodyPrefix = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetLogBodySuffix(v string) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.LogBodySuffix = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetLogSplit(v bool) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.LogSplit = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetLogSplitWords(v string) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.LogSplitWords = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetMaxBatchMB(v int64) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.MaxBatchMB = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetMaxBatchSize(v int64) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.MaxBatchSize = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetMaxRetry(v int64) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.MaxRetry = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetQueryParam(v map[string]*HttpDeliveryQueryParamValue) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.QueryParam = v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetStandardAuthOn(v bool) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.StandardAuthOn = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetStandardAuthParam(v *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.StandardAuthParam = v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) SetTransformTimeout(v int64) *CreateSiteDeliveryTaskRequestHttpDelivery {
	s.TransformTimeout = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDelivery) Validate() error {
	if s.StandardAuthParam != nil {
		if err := s.StandardAuthParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam struct {
	// The encryption timeout period.
	//
	// > The value must be greater than 0. We recommend that you set the value to 300 or greater.
	//
	// example:
	//
	// 300
	ExpiredTime *int32 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The private key.
	//
	// example:
	//
	// ***
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The URI path for standard authentication.
	//
	// example:
	//
	// v1/log/upload
	UrlPath *string `json:"UrlPath,omitempty" xml:"UrlPath,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) GetExpiredTime() *int32 {
	return s.ExpiredTime
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) GetUrlPath() *string {
	return s.UrlPath
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetExpiredTime(v int32) *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.ExpiredTime = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetPrivateKey(v string) *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.PrivateKey = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) SetUrlPath(v string) *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam {
	s.UrlPath = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestHttpDeliveryStandardAuthParam) Validate() error {
	return dara.Validate(s)
}

type CreateSiteDeliveryTaskRequestKafkaDelivery struct {
	// The load balancing method.
	//
	// example:
	//
	// kafka.LeastBytes
	Balancer *string `json:"Balancer,omitempty" xml:"Balancer,omitempty"`
	// The array of servers.
	Brokers []*string `json:"Brokers,omitempty" xml:"Brokers,omitempty" type:"Repeated"`
	// The compression method.
	//
	// example:
	//
	// lz4
	Compress *string `json:"Compress,omitempty" xml:"Compress,omitempty"`
	// The encryption method.
	//
	// example:
	//
	// plain
	MachanismType *string `json:"MachanismType,omitempty" xml:"MachanismType,omitempty"`
	// The password for encryption.
	//
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The Kafka topic.
	//
	// example:
	//
	// dqc_test2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// Specifies whether to enable user authentication.
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

func (s CreateSiteDeliveryTaskRequestKafkaDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestKafkaDelivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) GetBalancer() *string {
	return s.Balancer
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) GetBrokers() []*string {
	return s.Brokers
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) GetCompress() *string {
	return s.Compress
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) GetMachanismType() *string {
	return s.MachanismType
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) GetPassword() *string {
	return s.Password
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) GetTopic() *string {
	return s.Topic
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) GetUserAuth() *bool {
	return s.UserAuth
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) GetUserName() *string {
	return s.UserName
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetBalancer(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Balancer = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetBrokers(v []*string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Brokers = v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetCompress(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Compress = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetMachanismType(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.MachanismType = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetPassword(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Password = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetTopic(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.Topic = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetUserAuth(v bool) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.UserAuth = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) SetUserName(v string) *CreateSiteDeliveryTaskRequestKafkaDelivery {
	s.UserName = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestKafkaDelivery) Validate() error {
	return dara.Validate(s)
}

type CreateSiteDeliveryTaskRequestOssDelivery struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1234***
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The bucket name.
	//
	// example:
	//
	// test_rlog
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The prefix of the object key when logs are stored in the OSS bucket.
	//
	// example:
	//
	// test/
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// The region of the OSS bucket.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestOssDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestOssDelivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) GetAliuid() *string {
	return s.Aliuid
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) GetPrefixPath() *string {
	return s.PrefixPath
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) GetRegion() *string {
	return s.Region
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) SetAliuid(v string) *CreateSiteDeliveryTaskRequestOssDelivery {
	s.Aliuid = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) SetBucketName(v string) *CreateSiteDeliveryTaskRequestOssDelivery {
	s.BucketName = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) SetPrefixPath(v string) *CreateSiteDeliveryTaskRequestOssDelivery {
	s.PrefixPath = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) SetRegion(v string) *CreateSiteDeliveryTaskRequestOssDelivery {
	s.Region = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestOssDelivery) Validate() error {
	return dara.Validate(s)
}

type CreateSiteDeliveryTaskRequestS3Delivery struct {
	// The AccessKey ID of the Alibaba Cloud account or RAM user.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The bucket path.
	//
	// example:
	//
	// logriver-test/log
	BucketPath *string `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	// The endpoint of the server. This parameter is required when S3Cmpt is set to true.
	//
	// > For S3-compatible services, you must configure DNS resolution for a domain name that is spliced by the bucket and endpoint. For example, if Endpoint is set to example.com and Bucket is set to demo, the actual delivery address is demo.example.com.
	//
	// example:
	//
	// https://s3.oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The storage path prefix.
	//
	// example:
	//
	// logriver-test/log
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// The region where the service is located.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Specifies whether the service is S3-compatible.
	//
	// example:
	//
	// true
	S3Cmpt *bool `json:"S3Cmpt,omitempty" xml:"S3Cmpt,omitempty"`
	// The AccessKey secret of the S3 account.
	//
	// example:
	//
	// LDSIKh***
	SecretKey            *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	ServerSideEncryption *bool   `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	VertifyType          *string `json:"VertifyType,omitempty" xml:"VertifyType,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestS3Delivery) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestS3Delivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) GetBucketPath() *string {
	return s.BucketPath
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) GetPrefixPath() *string {
	return s.PrefixPath
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) GetRegion() *string {
	return s.Region
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) GetS3Cmpt() *bool {
	return s.S3Cmpt
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) GetSecretKey() *string {
	return s.SecretKey
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) GetServerSideEncryption() *bool {
	return s.ServerSideEncryption
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) GetVertifyType() *string {
	return s.VertifyType
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetAccessKey(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.AccessKey = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetBucketPath(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.BucketPath = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetEndpoint(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.Endpoint = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetPrefixPath(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.PrefixPath = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetRegion(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.Region = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetS3Cmpt(v bool) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.S3Cmpt = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetSecretKey(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.SecretKey = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetServerSideEncryption(v bool) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.ServerSideEncryption = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) SetVertifyType(v string) *CreateSiteDeliveryTaskRequestS3Delivery {
	s.VertifyType = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestS3Delivery) Validate() error {
	return dara.Validate(s)
}

type CreateSiteDeliveryTaskRequestSlsDelivery struct {
	// The name of the Logstore in Log Service.
	//
	// example:
	//
	// accesslog-test
	SLSLogStore *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	// The name of the Log Service project.
	//
	// example:
	//
	// dcdn-test20240417
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	// The region of the Log Service project.
	//
	// example:
	//
	// cn-hangzhou
	SLSRegion *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
}

func (s CreateSiteDeliveryTaskRequestSlsDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteDeliveryTaskRequestSlsDelivery) GoString() string {
	return s.String()
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) GetSLSLogStore() *string {
	return s.SLSLogStore
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) GetSLSProject() *string {
	return s.SLSProject
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) GetSLSRegion() *string {
	return s.SLSRegion
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) SetSLSLogStore(v string) *CreateSiteDeliveryTaskRequestSlsDelivery {
	s.SLSLogStore = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) SetSLSProject(v string) *CreateSiteDeliveryTaskRequestSlsDelivery {
	s.SLSProject = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) SetSLSRegion(v string) *CreateSiteDeliveryTaskRequestSlsDelivery {
	s.SLSRegion = &v
	return s
}

func (s *CreateSiteDeliveryTaskRequestSlsDelivery) Validate() error {
	return dara.Validate(s)
}
