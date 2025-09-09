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
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
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
