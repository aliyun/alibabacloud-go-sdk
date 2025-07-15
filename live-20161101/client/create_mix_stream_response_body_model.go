// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMixStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMixStreamId(v string) *CreateMixStreamResponseBody
	GetMixStreamId() *string
	SetRequestId(v string) *CreateMixStreamResponseBody
	GetRequestId() *string
}

type CreateMixStreamResponseBody struct {
	// The ID of the stream mixing task. You can specify this parameter in a request to delete the stream mixing task or query stream mixing tasks.
	//
	// example:
	//
	// 5b2a046e-74d7-385e-253f-8a5b87e4****
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0D715397-2E66-4AE1-694h-C546628AD145
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMixStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMixStreamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMixStreamResponseBody) GetMixStreamId() *string {
	return s.MixStreamId
}

func (s *CreateMixStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMixStreamResponseBody) SetMixStreamId(v string) *CreateMixStreamResponseBody {
	s.MixStreamId = &v
	return s
}

func (s *CreateMixStreamResponseBody) SetRequestId(v string) *CreateMixStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMixStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
