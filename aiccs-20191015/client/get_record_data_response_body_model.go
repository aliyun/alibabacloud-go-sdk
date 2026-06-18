// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRecordDataResponseBody
	GetCode() *string
	SetData(v *GetRecordDataResponseBodyData) *GetRecordDataResponseBody
	GetData() *GetRecordDataResponseBodyData
	SetHttpStatusCode(v int64) *GetRecordDataResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetRecordDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRecordDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRecordDataResponseBody
	GetSuccess() *bool
}

type GetRecordDataResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data list.
	Data *GetRecordDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRecordDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecordDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRecordDataResponseBody) GetData() *GetRecordDataResponseBodyData {
	return s.Data
}

func (s *GetRecordDataResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetRecordDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRecordDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecordDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRecordDataResponseBody) SetCode(v string) *GetRecordDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecordDataResponseBody) SetData(v *GetRecordDataResponseBodyData) *GetRecordDataResponseBody {
	s.Data = v
	return s
}

func (s *GetRecordDataResponseBody) SetHttpStatusCode(v int64) *GetRecordDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRecordDataResponseBody) SetMessage(v string) *GetRecordDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecordDataResponseBody) SetRequestId(v string) *GetRecordDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordDataResponseBody) SetSuccess(v bool) *GetRecordDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetRecordDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRecordDataResponseBodyData struct {
	// Session ID.
	//
	// example:
	//
	// 1001067****
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// Recording file URL.
	//
	// example:
	//
	// http://aliccrec-shvpc.oss-cn-shanghai.aliyuncs.com/accrec_tmp/10010679716-12-01-56.wav?***
	OssLink *string `json:"OssLink,omitempty" xml:"OssLink,omitempty"`
}

func (s GetRecordDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRecordDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecordDataResponseBodyData) GetAcid() *string {
	return s.Acid
}

func (s *GetRecordDataResponseBodyData) GetOssLink() *string {
	return s.OssLink
}

func (s *GetRecordDataResponseBodyData) SetAcid(v string) *GetRecordDataResponseBodyData {
	s.Acid = &v
	return s
}

func (s *GetRecordDataResponseBodyData) SetOssLink(v string) *GetRecordDataResponseBodyData {
	s.OssLink = &v
	return s
}

func (s *GetRecordDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
