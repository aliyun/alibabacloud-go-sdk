// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOrUpdateOssConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveOrUpdateOssConfigResponseBody
	GetCode() *string
	SetData(v *SaveOrUpdateOssConfigResponseBodyData) *SaveOrUpdateOssConfigResponseBody
	GetData() *SaveOrUpdateOssConfigResponseBodyData
	SetHttpStatusCode(v int32) *SaveOrUpdateOssConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveOrUpdateOssConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveOrUpdateOssConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveOrUpdateOssConfigResponseBody
	GetSuccess() *bool
}

type SaveOrUpdateOssConfigResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SaveOrUpdateOssConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveOrUpdateOssConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveOrUpdateOssConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveOrUpdateOssConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveOrUpdateOssConfigResponseBody) GetData() *SaveOrUpdateOssConfigResponseBodyData {
	return s.Data
}

func (s *SaveOrUpdateOssConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveOrUpdateOssConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveOrUpdateOssConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveOrUpdateOssConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveOrUpdateOssConfigResponseBody) SetCode(v string) *SaveOrUpdateOssConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SaveOrUpdateOssConfigResponseBody) SetData(v *SaveOrUpdateOssConfigResponseBodyData) *SaveOrUpdateOssConfigResponseBody {
	s.Data = v
	return s
}

func (s *SaveOrUpdateOssConfigResponseBody) SetHttpStatusCode(v int32) *SaveOrUpdateOssConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveOrUpdateOssConfigResponseBody) SetMessage(v string) *SaveOrUpdateOssConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SaveOrUpdateOssConfigResponseBody) SetRequestId(v string) *SaveOrUpdateOssConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveOrUpdateOssConfigResponseBody) SetSuccess(v bool) *SaveOrUpdateOssConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SaveOrUpdateOssConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveOrUpdateOssConfigResponseBodyData struct {
	// example:
	//
	// xxx
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 1
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
}

func (s SaveOrUpdateOssConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SaveOrUpdateOssConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *SaveOrUpdateOssConfigResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *SaveOrUpdateOssConfigResponseBodyData) GetEnable() *string {
	return s.Enable
}

func (s *SaveOrUpdateOssConfigResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SaveOrUpdateOssConfigResponseBodyData) SetBucketName(v string) *SaveOrUpdateOssConfigResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *SaveOrUpdateOssConfigResponseBodyData) SetEnable(v string) *SaveOrUpdateOssConfigResponseBodyData {
	s.Enable = &v
	return s
}

func (s *SaveOrUpdateOssConfigResponseBodyData) SetEndpoint(v string) *SaveOrUpdateOssConfigResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *SaveOrUpdateOssConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
