// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIImageInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAIImageInfosResponseBody
	GetRequestId() *string
}

type DeleteAIImageInfosResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FCDC80EA-363C-41*****B8-0DF14033D643
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIImageInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIImageInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIImageInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAIImageInfosResponseBody) SetRequestId(v string) *DeleteAIImageInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAIImageInfosResponseBody) Validate() error {
	return dara.Validate(s)
}
