// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteArtifactResponseBody
	GetRequestId() *string
}

type DeleteArtifactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1B3AD3CC-E938-5042-A771-7FD9A2FE03F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteArtifactResponseBody) SetRequestId(v string) *DeleteArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteArtifactResponseBody) Validate() error {
	return dara.Validate(s)
}
