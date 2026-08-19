// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomerLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCustomerLabelResponseBody
	GetCode() *string
	SetData(v *QueryCustomerLabelResponseBodyData) *QueryCustomerLabelResponseBody
	GetData() *QueryCustomerLabelResponseBodyData
	SetMessage(v string) *QueryCustomerLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCustomerLabelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCustomerLabelResponseBody
	GetSuccess() *bool
}

type QueryCustomerLabelResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryCustomerLabelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCustomerLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerLabelResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCustomerLabelResponseBody) GetData() *QueryCustomerLabelResponseBodyData {
	return s.Data
}

func (s *QueryCustomerLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCustomerLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCustomerLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCustomerLabelResponseBody) SetCode(v string) *QueryCustomerLabelResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomerLabelResponseBody) SetData(v *QueryCustomerLabelResponseBodyData) *QueryCustomerLabelResponseBody {
	s.Data = v
	return s
}

func (s *QueryCustomerLabelResponseBody) SetMessage(v string) *QueryCustomerLabelResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomerLabelResponseBody) SetRequestId(v string) *QueryCustomerLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomerLabelResponseBody) SetSuccess(v bool) *QueryCustomerLabelResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCustomerLabelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCustomerLabelResponseBodyData struct {
	CustomerLabel []*QueryCustomerLabelResponseBodyDataCustomerLabel `json:"CustomerLabel,omitempty" xml:"CustomerLabel,omitempty" type:"Repeated"`
}

func (s QueryCustomerLabelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerLabelResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCustomerLabelResponseBodyData) GetCustomerLabel() []*QueryCustomerLabelResponseBodyDataCustomerLabel {
	return s.CustomerLabel
}

func (s *QueryCustomerLabelResponseBodyData) SetCustomerLabel(v []*QueryCustomerLabelResponseBodyDataCustomerLabel) *QueryCustomerLabelResponseBodyData {
	s.CustomerLabel = v
	return s
}

func (s *QueryCustomerLabelResponseBodyData) Validate() error {
	if s.CustomerLabel != nil {
		for _, item := range s.CustomerLabel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCustomerLabelResponseBodyDataCustomerLabel struct {
	Creator        *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	EndTimeStr     *string `json:"EndTimeStr,omitempty" xml:"EndTimeStr,omitempty"`
	GmtCreatedStr  *string `json:"GmtCreatedStr,omitempty" xml:"GmtCreatedStr,omitempty"`
	GmtModifiedStr *string `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Label          *string `json:"Label,omitempty" xml:"Label,omitempty"`
	LabelSeries    *string `json:"LabelSeries,omitempty" xml:"LabelSeries,omitempty"`
	StartTimeStr   *string `json:"StartTimeStr,omitempty" xml:"StartTimeStr,omitempty"`
}

func (s QueryCustomerLabelResponseBodyDataCustomerLabel) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerLabelResponseBodyDataCustomerLabel) GoString() string {
	return s.String()
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) GetCreator() *string {
	return s.Creator
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) GetEndTimeStr() *string {
	return s.EndTimeStr
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) GetGmtCreatedStr() *string {
	return s.GmtCreatedStr
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) GetGmtModifiedStr() *string {
	return s.GmtModifiedStr
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) GetId() *string {
	return s.Id
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) GetLabel() *string {
	return s.Label
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) GetLabelSeries() *string {
	return s.LabelSeries
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) GetStartTimeStr() *string {
	return s.StartTimeStr
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) SetCreator(v string) *QueryCustomerLabelResponseBodyDataCustomerLabel {
	s.Creator = &v
	return s
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) SetEndTimeStr(v string) *QueryCustomerLabelResponseBodyDataCustomerLabel {
	s.EndTimeStr = &v
	return s
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) SetGmtCreatedStr(v string) *QueryCustomerLabelResponseBodyDataCustomerLabel {
	s.GmtCreatedStr = &v
	return s
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) SetGmtModifiedStr(v string) *QueryCustomerLabelResponseBodyDataCustomerLabel {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) SetId(v string) *QueryCustomerLabelResponseBodyDataCustomerLabel {
	s.Id = &v
	return s
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) SetLabel(v string) *QueryCustomerLabelResponseBodyDataCustomerLabel {
	s.Label = &v
	return s
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) SetLabelSeries(v string) *QueryCustomerLabelResponseBodyDataCustomerLabel {
	s.LabelSeries = &v
	return s
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) SetStartTimeStr(v string) *QueryCustomerLabelResponseBodyDataCustomerLabel {
	s.StartTimeStr = &v
	return s
}

func (s *QueryCustomerLabelResponseBodyDataCustomerLabel) Validate() error {
	return dara.Validate(s)
}
