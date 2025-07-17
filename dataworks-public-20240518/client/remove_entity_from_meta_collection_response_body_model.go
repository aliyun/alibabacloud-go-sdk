// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntityFromMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveEntityFromMetaCollectionResponseBody
	GetRequestId() *string
}

type RemoveEntityFromMetaCollectionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6D6CD444-DFA0-5180-9763-4A8730F2B382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveEntityFromMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntityFromMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntityFromMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveEntityFromMetaCollectionResponseBody) SetRequestId(v string) *RemoveEntityFromMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveEntityFromMetaCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
