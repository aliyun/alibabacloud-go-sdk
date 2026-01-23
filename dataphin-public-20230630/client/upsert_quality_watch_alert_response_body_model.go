// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityWatchAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpsertQualityWatchAlertResponseBody
	GetCode() *string
	SetData(v bool) *UpsertQualityWatchAlertResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpsertQualityWatchAlertResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpsertQualityWatchAlertResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpsertQualityWatchAlertResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpsertQualityWatchAlertResponseBody
	GetSuccess() *bool
}

type UpsertQualityWatchAlertResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpsertQualityWatchAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchAlertResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchAlertResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpsertQualityWatchAlertResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpsertQualityWatchAlertResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpsertQualityWatchAlertResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpsertQualityWatchAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertQualityWatchAlertResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpsertQualityWatchAlertResponseBody) SetCode(v string) *UpsertQualityWatchAlertResponseBody {
	s.Code = &v
	return s
}

func (s *UpsertQualityWatchAlertResponseBody) SetData(v bool) *UpsertQualityWatchAlertResponseBody {
	s.Data = &v
	return s
}

func (s *UpsertQualityWatchAlertResponseBody) SetHttpStatusCode(v int32) *UpsertQualityWatchAlertResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpsertQualityWatchAlertResponseBody) SetMessage(v string) *UpsertQualityWatchAlertResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertQualityWatchAlertResponseBody) SetRequestId(v string) *UpsertQualityWatchAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertQualityWatchAlertResponseBody) SetSuccess(v bool) *UpsertQualityWatchAlertResponseBody {
	s.Success = &v
	return s
}

func (s *UpsertQualityWatchAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
