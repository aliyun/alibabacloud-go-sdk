// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOptionValueForProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOptionValue(v string) *GetOptionValueForProjectResponseBody
	GetOptionValue() *string
	SetRequestId(v string) *GetOptionValueForProjectResponseBody
	GetRequestId() *string
}

type GetOptionValueForProjectResponseBody struct {
	// The data returned In the example, cuNumber is a custom key.
	//
	// example:
	//
	// {"cuNumber":"0"}
	OptionValue *string `json:"OptionValue,omitempty" xml:"OptionValue,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOptionValueForProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOptionValueForProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetOptionValueForProjectResponseBody) GetOptionValue() *string {
	return s.OptionValue
}

func (s *GetOptionValueForProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOptionValueForProjectResponseBody) SetOptionValue(v string) *GetOptionValueForProjectResponseBody {
	s.OptionValue = &v
	return s
}

func (s *GetOptionValueForProjectResponseBody) SetRequestId(v string) *GetOptionValueForProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOptionValueForProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
