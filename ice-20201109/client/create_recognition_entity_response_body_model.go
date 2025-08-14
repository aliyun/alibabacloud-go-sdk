// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecognitionEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v string) *CreateRecognitionEntityResponseBody
	GetEntityId() *string
	SetRequestId(v string) *CreateRecognitionEntityResponseBody
	GetRequestId() *string
}

type CreateRecognitionEntityResponseBody struct {
	// example:
	//
	// **************544cb84754************
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecognitionEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecognitionEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecognitionEntityResponseBody) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateRecognitionEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecognitionEntityResponseBody) SetEntityId(v string) *CreateRecognitionEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *CreateRecognitionEntityResponseBody) SetRequestId(v string) *CreateRecognitionEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecognitionEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
