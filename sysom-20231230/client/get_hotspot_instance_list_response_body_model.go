// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotspotInstanceListResponseBody
	GetCode() *string
	SetData(v *GetHotspotInstanceListResponseBodyData) *GetHotspotInstanceListResponseBody
	GetData() *GetHotspotInstanceListResponseBodyData
	SetMessage(v string) *GetHotspotInstanceListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotspotInstanceListResponseBody
	GetRequestId() *string
}

type GetHotspotInstanceListResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotspotInstanceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHotspotInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotInstanceListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotspotInstanceListResponseBody) GetData() *GetHotspotInstanceListResponseBodyData {
	return s.Data
}

func (s *GetHotspotInstanceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotspotInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotspotInstanceListResponseBody) SetCode(v string) *GetHotspotInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotInstanceListResponseBody) SetData(v *GetHotspotInstanceListResponseBodyData) *GetHotspotInstanceListResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotInstanceListResponseBody) SetMessage(v string) *GetHotspotInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotInstanceListResponseBody) SetRequestId(v string) *GetHotspotInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotInstanceListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotspotInstanceListResponseBodyData struct {
	Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotInstanceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotInstanceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotInstanceListResponseBodyData) GetColumns() []*string {
	return s.Columns
}

func (s *GetHotspotInstanceListResponseBodyData) GetValues() []*string {
	return s.Values
}

func (s *GetHotspotInstanceListResponseBodyData) SetColumns(v []*string) *GetHotspotInstanceListResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotspotInstanceListResponseBodyData) SetValues(v []*string) *GetHotspotInstanceListResponseBodyData {
	s.Values = v
	return s
}

func (s *GetHotspotInstanceListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
