// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavedQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *CreateSavedQueryResponseBody
	GetQueryId() *string
	SetRequestId(v string) *CreateSavedQueryResponseBody
	GetRequestId() *string
}

type CreateSavedQueryResponseBody struct {
	// The template ID.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFA806B9-7F36-55AB-8B7A-D680C2C5EE57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSavedQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSavedQueryResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *CreateSavedQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSavedQueryResponseBody) SetQueryId(v string) *CreateSavedQueryResponseBody {
	s.QueryId = &v
	return s
}

func (s *CreateSavedQueryResponseBody) SetRequestId(v string) *CreateSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSavedQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
