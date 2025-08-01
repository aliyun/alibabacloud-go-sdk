// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotPidListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotspotPidListResponseBody
	GetCode() *string
	SetData(v *GetHotspotPidListResponseBodyData) *GetHotspotPidListResponseBody
	GetData() *GetHotspotPidListResponseBodyData
	SetMessage(v string) *GetHotspotPidListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotspotPidListResponseBody
	GetRequestId() *string
}

type GetHotspotPidListResponseBody struct {
	// example:
	//
	// SysomOpenAPI.InvalidParameter
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotspotPidListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHotspotPidListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotPidListResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotPidListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotspotPidListResponseBody) GetData() *GetHotspotPidListResponseBodyData {
	return s.Data
}

func (s *GetHotspotPidListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotspotPidListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotspotPidListResponseBody) SetCode(v string) *GetHotspotPidListResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotPidListResponseBody) SetData(v *GetHotspotPidListResponseBodyData) *GetHotspotPidListResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotPidListResponseBody) SetMessage(v string) *GetHotspotPidListResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotPidListResponseBody) SetRequestId(v string) *GetHotspotPidListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotPidListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHotspotPidListResponseBodyData struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotPidListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotPidListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotPidListResponseBodyData) GetColumns() []*string {
	return s.Columns
}

func (s *GetHotspotPidListResponseBodyData) GetValues() [][]*string {
	return s.Values
}

func (s *GetHotspotPidListResponseBodyData) SetColumns(v []*string) *GetHotspotPidListResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotspotPidListResponseBodyData) SetValues(v [][]*string) *GetHotspotPidListResponseBodyData {
	s.Values = v
	return s
}

func (s *GetHotspotPidListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
