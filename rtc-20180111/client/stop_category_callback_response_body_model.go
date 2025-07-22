// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCategoryCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopCategoryCallbackResponseBody
	GetRequestId() *string
}

type StopCategoryCallbackResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2DCE8D7E-BE3B-54AB-8DAC-32F34BED0763
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopCategoryCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCategoryCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *StopCategoryCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCategoryCallbackResponseBody) SetRequestId(v string) *StopCategoryCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCategoryCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
