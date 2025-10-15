// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDialogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *CreateDialogResponseBody
	GetCost() *int64
	SetData(v *CreateDialogResponseBodyData) *CreateDialogResponseBody
	GetData() *CreateDialogResponseBodyData
	SetDataType(v string) *CreateDialogResponseBody
	GetDataType() *string
	SetErrCode(v string) *CreateDialogResponseBody
	GetErrCode() *string
	SetMessage(v string) *CreateDialogResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDialogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDialogResponseBody
	GetSuccess() *bool
	SetTime(v string) *CreateDialogResponseBody
	GetTime() *string
}

type CreateDialogResponseBody struct {
	// example:
	//
	// null
	Cost *int64                        `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *CreateDialogResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 003D019A-1BB3-53EC-A0D2-CE76DA5D73B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateDialogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDialogResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *CreateDialogResponseBody) GetData() *CreateDialogResponseBodyData {
	return s.Data
}

func (s *CreateDialogResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *CreateDialogResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateDialogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDialogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDialogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDialogResponseBody) GetTime() *string {
	return s.Time
}

func (s *CreateDialogResponseBody) SetCost(v int64) *CreateDialogResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateDialogResponseBody) SetData(v *CreateDialogResponseBodyData) *CreateDialogResponseBody {
	s.Data = v
	return s
}

func (s *CreateDialogResponseBody) SetDataType(v string) *CreateDialogResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateDialogResponseBody) SetErrCode(v string) *CreateDialogResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDialogResponseBody) SetMessage(v string) *CreateDialogResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDialogResponseBody) SetRequestId(v string) *CreateDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDialogResponseBody) SetSuccess(v bool) *CreateDialogResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDialogResponseBody) SetTime(v string) *CreateDialogResponseBody {
	s.Time = &v
	return s
}

func (s *CreateDialogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDialogResponseBodyData struct {
	OpeningRemarks *string `json:"openingRemarks,omitempty" xml:"openingRemarks,omitempty"`
	// example:
	//
	// 1728545917713234
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CreateDialogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDialogResponseBodyData) GetOpeningRemarks() *string {
	return s.OpeningRemarks
}

func (s *CreateDialogResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateDialogResponseBodyData) SetOpeningRemarks(v string) *CreateDialogResponseBodyData {
	s.OpeningRemarks = &v
	return s
}

func (s *CreateDialogResponseBodyData) SetSessionId(v string) *CreateDialogResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *CreateDialogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
