// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveAIStudioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLiveAIStudioResponseBody
	GetRequestId() *string
	SetStudioId(v string) *CreateLiveAIStudioResponseBody
	GetStudioId() *string
}

type CreateLiveAIStudioResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 17D7526C-69AD-5761-8037-071C27358345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the virtual studio template.
	//
	// example:
	//
	// 369ced1f-c33a-49e5-91da-bdaae3d6c1c2
	StudioId *string `json:"StudioId,omitempty" xml:"StudioId,omitempty"`
}

func (s CreateLiveAIStudioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveAIStudioResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveAIStudioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveAIStudioResponseBody) GetStudioId() *string {
	return s.StudioId
}

func (s *CreateLiveAIStudioResponseBody) SetRequestId(v string) *CreateLiveAIStudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveAIStudioResponseBody) SetStudioId(v string) *CreateLiveAIStudioResponseBody {
	s.StudioId = &v
	return s
}

func (s *CreateLiveAIStudioResponseBody) Validate() error {
	return dara.Validate(s)
}
