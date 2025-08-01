// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotSpotUniqListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotSpotUniqListResponseBody
	GetCode() *string
	SetData(v *GetHotSpotUniqListResponseBodyData) *GetHotSpotUniqListResponseBody
	GetData() *GetHotSpotUniqListResponseBodyData
	SetMessage(v string) *GetHotSpotUniqListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotSpotUniqListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetHotSpotUniqListResponseBody
	GetSuccess() *string
}

type GetHotSpotUniqListResponseBody struct {
	// example:
	//
	// Success
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotSpotUniqListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHotSpotUniqListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotSpotUniqListResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotSpotUniqListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotSpotUniqListResponseBody) GetData() *GetHotSpotUniqListResponseBodyData {
	return s.Data
}

func (s *GetHotSpotUniqListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotSpotUniqListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotSpotUniqListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetHotSpotUniqListResponseBody) SetCode(v string) *GetHotSpotUniqListResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotSpotUniqListResponseBody) SetData(v *GetHotSpotUniqListResponseBodyData) *GetHotSpotUniqListResponseBody {
	s.Data = v
	return s
}

func (s *GetHotSpotUniqListResponseBody) SetMessage(v string) *GetHotSpotUniqListResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotSpotUniqListResponseBody) SetRequestId(v string) *GetHotSpotUniqListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotSpotUniqListResponseBody) SetSuccess(v string) *GetHotSpotUniqListResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotSpotUniqListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHotSpotUniqListResponseBodyData struct {
	Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotSpotUniqListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotSpotUniqListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotSpotUniqListResponseBodyData) GetColumns() []*string {
	return s.Columns
}

func (s *GetHotSpotUniqListResponseBodyData) GetValues() []*string {
	return s.Values
}

func (s *GetHotSpotUniqListResponseBodyData) SetColumns(v []*string) *GetHotSpotUniqListResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotSpotUniqListResponseBodyData) SetValues(v []*string) *GetHotSpotUniqListResponseBodyData {
	s.Values = v
	return s
}

func (s *GetHotSpotUniqListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
