// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountLabelList(v []*GetAccountLabelResponseBodyAccountLabelList) *GetAccountLabelResponseBody
	GetAccountLabelList() []*GetAccountLabelResponseBodyAccountLabelList
	SetRequestId(v string) *GetAccountLabelResponseBody
	GetRequestId() *string
}

type GetAccountLabelResponseBody struct {
	// The tag list.
	AccountLabelList []*GetAccountLabelResponseBodyAccountLabelList `json:"AccountLabelList,omitempty" xml:"AccountLabelList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountLabelResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountLabelResponseBody) GetAccountLabelList() []*GetAccountLabelResponseBodyAccountLabelList {
	return s.AccountLabelList
}

func (s *GetAccountLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountLabelResponseBody) SetAccountLabelList(v []*GetAccountLabelResponseBodyAccountLabelList) *GetAccountLabelResponseBody {
	s.AccountLabelList = v
	return s
}

func (s *GetAccountLabelResponseBody) SetRequestId(v string) *GetAccountLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountLabelResponseBody) Validate() error {
	if s.AccountLabelList != nil {
		for _, item := range s.AccountLabelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAccountLabelResponseBodyAccountLabelList struct {
	// The tag information.
	//
	// example:
	//
	// SasStep
	LabelSeries *string `json:"LabelSeries,omitempty" xml:"LabelSeries,omitempty"`
	// Indicates whether the tag is valid.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	LabelStatus *bool `json:"LabelStatus,omitempty" xml:"LabelStatus,omitempty"`
}

func (s GetAccountLabelResponseBodyAccountLabelList) String() string {
	return dara.Prettify(s)
}

func (s GetAccountLabelResponseBodyAccountLabelList) GoString() string {
	return s.String()
}

func (s *GetAccountLabelResponseBodyAccountLabelList) GetLabelSeries() *string {
	return s.LabelSeries
}

func (s *GetAccountLabelResponseBodyAccountLabelList) GetLabelStatus() *bool {
	return s.LabelStatus
}

func (s *GetAccountLabelResponseBodyAccountLabelList) SetLabelSeries(v string) *GetAccountLabelResponseBodyAccountLabelList {
	s.LabelSeries = &v
	return s
}

func (s *GetAccountLabelResponseBodyAccountLabelList) SetLabelStatus(v bool) *GetAccountLabelResponseBodyAccountLabelList {
	s.LabelStatus = &v
	return s
}

func (s *GetAccountLabelResponseBodyAccountLabelList) Validate() error {
	return dara.Validate(s)
}
