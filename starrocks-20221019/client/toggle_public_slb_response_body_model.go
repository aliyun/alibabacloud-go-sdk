// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTogglePublicSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *TogglePublicSlbResponseBody
	GetData() *bool
	SetErrCode(v string) *TogglePublicSlbResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *TogglePublicSlbResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *TogglePublicSlbResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *TogglePublicSlbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TogglePublicSlbResponseBody
	GetSuccess() *bool
}

type TogglePublicSlbResponseBody struct {
	// example:
	//
	// 24151320976****
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TogglePublicSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TogglePublicSlbResponseBody) GoString() string {
	return s.String()
}

func (s *TogglePublicSlbResponseBody) GetData() *bool {
	return s.Data
}

func (s *TogglePublicSlbResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *TogglePublicSlbResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *TogglePublicSlbResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TogglePublicSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TogglePublicSlbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TogglePublicSlbResponseBody) SetData(v bool) *TogglePublicSlbResponseBody {
	s.Data = &v
	return s
}

func (s *TogglePublicSlbResponseBody) SetErrCode(v string) *TogglePublicSlbResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TogglePublicSlbResponseBody) SetErrMessage(v string) *TogglePublicSlbResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *TogglePublicSlbResponseBody) SetHttpStatusCode(v int32) *TogglePublicSlbResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TogglePublicSlbResponseBody) SetRequestId(v string) *TogglePublicSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *TogglePublicSlbResponseBody) SetSuccess(v bool) *TogglePublicSlbResponseBody {
	s.Success = &v
	return s
}

func (s *TogglePublicSlbResponseBody) Validate() error {
	return dara.Validate(s)
}
