// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomerLabelByConfigGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCustomerLabelByConfigGroupResponseBody
	GetCode() *string
	SetData(v []*QueryCustomerLabelByConfigGroupResponseBodyData) *QueryCustomerLabelByConfigGroupResponseBody
	GetData() []*QueryCustomerLabelByConfigGroupResponseBodyData
	SetMessage(v string) *QueryCustomerLabelByConfigGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCustomerLabelByConfigGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCustomerLabelByConfigGroupResponseBody
	GetSuccess() *bool
}

type QueryCustomerLabelByConfigGroupResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryCustomerLabelByConfigGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCustomerLabelByConfigGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerLabelByConfigGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) GetData() []*QueryCustomerLabelByConfigGroupResponseBodyData {
	return s.Data
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) SetCode(v string) *QueryCustomerLabelByConfigGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) SetData(v []*QueryCustomerLabelByConfigGroupResponseBodyData) *QueryCustomerLabelByConfigGroupResponseBody {
	s.Data = v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) SetMessage(v string) *QueryCustomerLabelByConfigGroupResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) SetRequestId(v string) *QueryCustomerLabelByConfigGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) SetSuccess(v bool) *QueryCustomerLabelByConfigGroupResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCustomerLabelByConfigGroupResponseBodyData struct {
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	LabelSeries *string `json:"LabelSeries,omitempty" xml:"LabelSeries,omitempty"`
}

func (s QueryCustomerLabelByConfigGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerLabelByConfigGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCustomerLabelByConfigGroupResponseBodyData) GetLabel() *string {
	return s.Label
}

func (s *QueryCustomerLabelByConfigGroupResponseBodyData) GetLabelSeries() *string {
	return s.LabelSeries
}

func (s *QueryCustomerLabelByConfigGroupResponseBodyData) SetLabel(v string) *QueryCustomerLabelByConfigGroupResponseBodyData {
	s.Label = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponseBodyData) SetLabelSeries(v string) *QueryCustomerLabelByConfigGroupResponseBodyData {
	s.LabelSeries = &v
	return s
}

func (s *QueryCustomerLabelByConfigGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
