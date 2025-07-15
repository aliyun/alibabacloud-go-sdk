// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAIStudioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveAIStudioResponseBody
	GetRequestId() *string
}

type DeleteLiveAIStudioResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40A4F36D-A7CC-473A-88E7-154F92242566
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveAIStudioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAIStudioResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveAIStudioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveAIStudioResponseBody) SetRequestId(v string) *DeleteLiveAIStudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveAIStudioResponseBody) Validate() error {
	return dara.Validate(s)
}
