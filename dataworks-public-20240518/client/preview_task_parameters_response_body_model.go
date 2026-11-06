// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewTaskParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*PreviewTaskParametersResponseBodyItems) *PreviewTaskParametersResponseBody
	GetItems() []*PreviewTaskParametersResponseBodyItems
	SetRequestId(v string) *PreviewTaskParametersResponseBody
	GetRequestId() *string
}

type PreviewTaskParametersResponseBody struct {
	// The parameter preview results for each cycle.
	Items []*PreviewTaskParametersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreviewTaskParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewTaskParametersResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewTaskParametersResponseBody) GetItems() []*PreviewTaskParametersResponseBodyItems {
	return s.Items
}

func (s *PreviewTaskParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewTaskParametersResponseBody) SetItems(v []*PreviewTaskParametersResponseBodyItems) *PreviewTaskParametersResponseBody {
	s.Items = v
	return s
}

func (s *PreviewTaskParametersResponseBody) SetRequestId(v string) *PreviewTaskParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewTaskParametersResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PreviewTaskParametersResponseBodyItems struct {
	// The business date of the cycle. The value is a 13-digit UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1792425600000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The error message returned when the parsing for the cycle fails. This parameter is returned when Success is false and is not returned when Success is true.
	//
	// example:
	//
	// Failed to parse scheduling parameters. Check the parameter settings.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The complete script parameter string parsed for the cycle. Multiple parameters are separated by spaces. This parameter is returned when Success is true.
	//
	// example:
	//
	// para1=20261020 para2=00
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Specifies whether the scheduling parameters for the cycle are parsed successfully. Valid values:
	//
	// - true: The parsing is successful.
	//
	// - false: The parsing fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The simulated scheduled trigger time. The value is a 13-digit UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1792512000000
	TriggerTime *int64 `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
}

func (s PreviewTaskParametersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s PreviewTaskParametersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *PreviewTaskParametersResponseBodyItems) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *PreviewTaskParametersResponseBodyItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PreviewTaskParametersResponseBodyItems) GetParameters() *string {
	return s.Parameters
}

func (s *PreviewTaskParametersResponseBodyItems) GetSuccess() *bool {
	return s.Success
}

func (s *PreviewTaskParametersResponseBodyItems) GetTriggerTime() *int64 {
	return s.TriggerTime
}

func (s *PreviewTaskParametersResponseBodyItems) SetBizdate(v int64) *PreviewTaskParametersResponseBodyItems {
	s.Bizdate = &v
	return s
}

func (s *PreviewTaskParametersResponseBodyItems) SetErrorMessage(v string) *PreviewTaskParametersResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *PreviewTaskParametersResponseBodyItems) SetParameters(v string) *PreviewTaskParametersResponseBodyItems {
	s.Parameters = &v
	return s
}

func (s *PreviewTaskParametersResponseBodyItems) SetSuccess(v bool) *PreviewTaskParametersResponseBodyItems {
	s.Success = &v
	return s
}

func (s *PreviewTaskParametersResponseBodyItems) SetTriggerTime(v int64) *PreviewTaskParametersResponseBodyItems {
	s.TriggerTime = &v
	return s
}

func (s *PreviewTaskParametersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
