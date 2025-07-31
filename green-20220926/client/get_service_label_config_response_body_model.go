// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLabelConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []interface{}) *GetServiceLabelConfigResponseBody
	GetData() []interface{}
	SetRequestId(v string) *GetServiceLabelConfigResponseBody
	GetRequestId() *string
}

type GetServiceLabelConfigResponseBody struct {
	Data []interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceLabelConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLabelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLabelConfigResponseBody) GetData() []interface{} {
	return s.Data
}

func (s *GetServiceLabelConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceLabelConfigResponseBody) SetData(v []interface{}) *GetServiceLabelConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceLabelConfigResponseBody) SetRequestId(v string) *GetServiceLabelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceLabelConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
