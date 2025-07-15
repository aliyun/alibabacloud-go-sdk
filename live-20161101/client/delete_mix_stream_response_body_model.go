// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMixStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMixStreamId(v string) *DeleteMixStreamResponseBody
	GetMixStreamId() *string
	SetRequestId(v string) *DeleteMixStreamResponseBody
	GetRequestId() *string
}

type DeleteMixStreamResponseBody struct {
	// The ID of the stream mixing task.
	//
	// example:
	//
	// 749b7594-86d6-37b1-513b-e1e19845****
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BE9407FF-F897-4DBD-338D-98A750AD805F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMixStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMixStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMixStreamResponseBody) GetMixStreamId() *string {
	return s.MixStreamId
}

func (s *DeleteMixStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMixStreamResponseBody) SetMixStreamId(v string) *DeleteMixStreamResponseBody {
	s.MixStreamId = &v
	return s
}

func (s *DeleteMixStreamResponseBody) SetRequestId(v string) *DeleteMixStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMixStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
