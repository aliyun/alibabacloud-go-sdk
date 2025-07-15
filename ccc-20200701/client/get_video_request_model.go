// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetVideoRequest
	GetContactId() *string
	SetInstanceId(v string) *GetVideoRequest
	GetInstanceId() *string
}

type GetVideoRequest struct {
	// This parameter is required.
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRequest) GoString() string {
	return s.String()
}

func (s *GetVideoRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetVideoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVideoRequest) SetContactId(v string) *GetVideoRequest {
	s.ContactId = &v
	return s
}

func (s *GetVideoRequest) SetInstanceId(v string) *GetVideoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVideoRequest) Validate() error {
	return dara.Validate(s)
}
