// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmsQualificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteSmsQualificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteSmsQualificationResponseBody
	GetCode() *string
	SetData(v bool) *DeleteSmsQualificationResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteSmsQualificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSmsQualificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSmsQualificationResponseBody
	GetSuccess() *bool
}

type DeleteSmsQualificationResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 25D5AFDE-8EBC-132E-8909-1FDC071DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSmsQualificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmsQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmsQualificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteSmsQualificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSmsQualificationResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteSmsQualificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSmsQualificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSmsQualificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSmsQualificationResponseBody) SetAccessDeniedDetail(v string) *DeleteSmsQualificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteSmsQualificationResponseBody) SetCode(v string) *DeleteSmsQualificationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSmsQualificationResponseBody) SetData(v bool) *DeleteSmsQualificationResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSmsQualificationResponseBody) SetMessage(v string) *DeleteSmsQualificationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSmsQualificationResponseBody) SetRequestId(v string) *DeleteSmsQualificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmsQualificationResponseBody) SetSuccess(v bool) *DeleteSmsQualificationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSmsQualificationResponseBody) Validate() error {
	return dara.Validate(s)
}
