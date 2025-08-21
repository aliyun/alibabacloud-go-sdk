// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRolloverDataStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RolloverDataStreamResponseBody
	GetRequestId() *string
	SetResult(v bool) *RolloverDataStreamResponseBody
	GetResult() *bool
}

type RolloverDataStreamResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RolloverDataStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RolloverDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *RolloverDataStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RolloverDataStreamResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RolloverDataStreamResponseBody) SetRequestId(v string) *RolloverDataStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *RolloverDataStreamResponseBody) SetResult(v bool) *RolloverDataStreamResponseBody {
	s.Result = &v
	return s
}

func (s *RolloverDataStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
