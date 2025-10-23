// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPcrInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPcrInfoResponseBodyData) *GetPcrInfoResponseBody
	GetData() *GetPcrInfoResponseBodyData
	SetRequestId(v string) *GetPcrInfoResponseBody
	GetRequestId() *string
}

type GetPcrInfoResponseBody struct {
	// The response parameters.
	Data *GetPcrInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 4A0AEC56-5C9A-5D47-93DF-7227836FFF82
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPcrInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPcrInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPcrInfoResponseBody) GetData() *GetPcrInfoResponseBodyData {
	return s.Data
}

func (s *GetPcrInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPcrInfoResponseBody) SetData(v *GetPcrInfoResponseBodyData) *GetPcrInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetPcrInfoResponseBody) SetRequestId(v string) *GetPcrInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPcrInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPcrInfoResponseBodyData struct {
	// The timestamp when the report was created. The timestamp is in milliseconds.
	//
	// example:
	//
	// 1709109790532
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Report name
	//
	// example:
	//
	// report name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Download url link.
	//
	// example:
	//
	// https://energy.alibabacloud.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetPcrInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPcrInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPcrInfoResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPcrInfoResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetPcrInfoResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetPcrInfoResponseBodyData) SetCreateTime(v string) *GetPcrInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetPcrInfoResponseBodyData) SetName(v string) *GetPcrInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPcrInfoResponseBodyData) SetUrl(v string) *GetPcrInfoResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetPcrInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
