// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeQualityProjectStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeQualityProjectStatusResponseBody
	GetCode() *string
	SetData(v string) *ChangeQualityProjectStatusResponseBody
	GetData() *string
	SetMessage(v string) *ChangeQualityProjectStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeQualityProjectStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeQualityProjectStatusResponseBody
	GetSuccess() *bool
}

type ChangeQualityProjectStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeQualityProjectStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeQualityProjectStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeQualityProjectStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeQualityProjectStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *ChangeQualityProjectStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeQualityProjectStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeQualityProjectStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeQualityProjectStatusResponseBody) SetCode(v string) *ChangeQualityProjectStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeQualityProjectStatusResponseBody) SetData(v string) *ChangeQualityProjectStatusResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeQualityProjectStatusResponseBody) SetMessage(v string) *ChangeQualityProjectStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeQualityProjectStatusResponseBody) SetRequestId(v string) *ChangeQualityProjectStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeQualityProjectStatusResponseBody) SetSuccess(v bool) *ChangeQualityProjectStatusResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeQualityProjectStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
