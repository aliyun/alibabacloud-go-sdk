// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDataStreamResponseBody
	GetRequestId() *string
	SetResult(v *CreateDataStreamResponseBodyResult) *CreateDataStreamResponseBody
	GetResult() *CreateDataStreamResponseBodyResult
}

type CreateDataStreamResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateDataStreamResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateDataStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataStreamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataStreamResponseBody) GetResult() *CreateDataStreamResponseBodyResult {
	return s.Result
}

func (s *CreateDataStreamResponseBody) SetRequestId(v string) *CreateDataStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataStreamResponseBody) SetResult(v *CreateDataStreamResponseBodyResult) *CreateDataStreamResponseBody {
	s.Result = v
	return s
}

func (s *CreateDataStreamResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDataStreamResponseBodyResult struct {
	// example:
	//
	// ds-
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateDataStreamResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateDataStreamResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDataStreamResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateDataStreamResponseBodyResult) SetName(v string) *CreateDataStreamResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateDataStreamResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
