// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilterConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AddFilterConfigsResponseBodyData) *AddFilterConfigsResponseBody
	GetData() []*AddFilterConfigsResponseBodyData
	SetRequestId(v string) *AddFilterConfigsResponseBody
	GetRequestId() *string
}

type AddFilterConfigsResponseBody struct {
	Data      []*AddFilterConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFilterConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFilterConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *AddFilterConfigsResponseBody) GetData() []*AddFilterConfigsResponseBodyData {
	return s.Data
}

func (s *AddFilterConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFilterConfigsResponseBody) SetData(v []*AddFilterConfigsResponseBodyData) *AddFilterConfigsResponseBody {
	s.Data = v
	return s
}

func (s *AddFilterConfigsResponseBody) SetRequestId(v string) *AddFilterConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFilterConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddFilterConfigsResponseBodyData struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Uuid    *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s AddFilterConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddFilterConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFilterConfigsResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *AddFilterConfigsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AddFilterConfigsResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *AddFilterConfigsResponseBodyData) SetMessage(v string) *AddFilterConfigsResponseBodyData {
	s.Message = &v
	return s
}

func (s *AddFilterConfigsResponseBodyData) SetSuccess(v bool) *AddFilterConfigsResponseBodyData {
	s.Success = &v
	return s
}

func (s *AddFilterConfigsResponseBodyData) SetUuid(v string) *AddFilterConfigsResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *AddFilterConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
