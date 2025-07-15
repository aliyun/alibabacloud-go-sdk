// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMixStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMixStreamId(v string) *UpdateMixStreamResponseBody
	GetMixStreamId() *string
	SetRequestId(v string) *UpdateMixStreamResponseBody
	GetRequestId() *string
}

type UpdateMixStreamResponseBody struct {
	// The ID of the stream mixing task. You can specify this parameter in a request to delete the stream mixing task or query stream mixing tasks.
	//
	// example:
	//
	// 5b2a046e-74d7-385e-d2d7-8a5b87e4****
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A8CDDFF-0121-4ABB-DA60-AEF095A8W34F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMixStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMixStreamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMixStreamResponseBody) GetMixStreamId() *string {
	return s.MixStreamId
}

func (s *UpdateMixStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMixStreamResponseBody) SetMixStreamId(v string) *UpdateMixStreamResponseBody {
	s.MixStreamId = &v
	return s
}

func (s *UpdateMixStreamResponseBody) SetRequestId(v string) *UpdateMixStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMixStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
