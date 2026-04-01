// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaSsoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKibanaSsoResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateKibanaSsoResponseBody
	GetResult() *bool
}

type UpdateKibanaSsoResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C82758DD-282F-4D48-934F-92170A33****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateKibanaSsoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaSsoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKibanaSsoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKibanaSsoResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateKibanaSsoResponseBody) SetRequestId(v string) *UpdateKibanaSsoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKibanaSsoResponseBody) SetResult(v bool) *UpdateKibanaSsoResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateKibanaSsoResponseBody) Validate() error {
	return dara.Validate(s)
}
