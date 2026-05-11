// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAsrConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyAsrConfigResponseBody
	GetCode() *string
	SetData(v *ModifyAsrConfigResponseBodyData) *ModifyAsrConfigResponseBody
	GetData() *ModifyAsrConfigResponseBodyData
	SetErrorMsg(v string) *ModifyAsrConfigResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *ModifyAsrConfigResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyAsrConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAsrConfigResponseBody
	GetSuccess() *bool
}

type ModifyAsrConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ModifyAsrConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Not Found
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAsrConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAsrConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAsrConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAsrConfigResponseBody) GetData() *ModifyAsrConfigResponseBodyData {
	return s.Data
}

func (s *ModifyAsrConfigResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ModifyAsrConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyAsrConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAsrConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAsrConfigResponseBody) SetCode(v string) *ModifyAsrConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetData(v *ModifyAsrConfigResponseBodyData) *ModifyAsrConfigResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetErrorMsg(v string) *ModifyAsrConfigResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetHttpStatusCode(v int32) *ModifyAsrConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetRequestId(v string) *ModifyAsrConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAsrConfigResponseBody) SetSuccess(v bool) *ModifyAsrConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAsrConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAsrConfigResponseBodyData struct {
	AffectedRows *int32 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
}

func (s ModifyAsrConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyAsrConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyAsrConfigResponseBodyData) GetAffectedRows() *int32 {
	return s.AffectedRows
}

func (s *ModifyAsrConfigResponseBodyData) SetAffectedRows(v int32) *ModifyAsrConfigResponseBodyData {
	s.AffectedRows = &v
	return s
}

func (s *ModifyAsrConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
