// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeHoloWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ResumeHoloWarehouseResponseBody
	GetData() *bool
	SetRequestId(v string) *ResumeHoloWarehouseResponseBody
	GetRequestId() *string
}

type ResumeHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeHoloWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeHoloWarehouseResponseBody) GetData() *bool {
	return s.Data
}

func (s *ResumeHoloWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeHoloWarehouseResponseBody) SetData(v bool) *ResumeHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeHoloWarehouseResponseBody) SetRequestId(v string) *ResumeHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeHoloWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}
