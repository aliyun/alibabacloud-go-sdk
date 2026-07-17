// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeEditingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEditingProjectId(v string) *CreateYikeEditingProjectResponseBody
	GetEditingProjectId() *string
	SetRequestId(v string) *CreateYikeEditingProjectResponseBody
	GetRequestId() *string
}

type CreateYikeEditingProjectResponseBody struct {
	// example:
	//
	// ***545bc38a94aa9840c89aff017b***
	EditingProjectId *string `json:"EditingProjectId,omitempty" xml:"EditingProjectId,omitempty"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateYikeEditingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateYikeEditingProjectResponseBody) GetEditingProjectId() *string {
	return s.EditingProjectId
}

func (s *CreateYikeEditingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateYikeEditingProjectResponseBody) SetEditingProjectId(v string) *CreateYikeEditingProjectResponseBody {
	s.EditingProjectId = &v
	return s
}

func (s *CreateYikeEditingProjectResponseBody) SetRequestId(v string) *CreateYikeEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateYikeEditingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
