// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilterConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteFilterConfigsResponseBodyData) *DeleteFilterConfigsResponseBody
	GetData() *DeleteFilterConfigsResponseBodyData
	SetRequestId(v string) *DeleteFilterConfigsResponseBody
	GetRequestId() *string
}

type DeleteFilterConfigsResponseBody struct {
	Data      *DeleteFilterConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFilterConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilterConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFilterConfigsResponseBody) GetData() *DeleteFilterConfigsResponseBodyData {
	return s.Data
}

func (s *DeleteFilterConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFilterConfigsResponseBody) SetData(v *DeleteFilterConfigsResponseBodyData) *DeleteFilterConfigsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteFilterConfigsResponseBody) SetRequestId(v string) *DeleteFilterConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFilterConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteFilterConfigsResponseBodyData struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFilterConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilterConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteFilterConfigsResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteFilterConfigsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFilterConfigsResponseBodyData) SetMessage(v string) *DeleteFilterConfigsResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteFilterConfigsResponseBodyData) SetSuccess(v bool) *DeleteFilterConfigsResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteFilterConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
