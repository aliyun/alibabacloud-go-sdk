// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTurnCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTurnCredentialsRequest
	GetInstanceId() *string
	SetUserId(v string) *GetTurnCredentialsRequest
	GetUserId() *string
}

type GetTurnCredentialsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTurnCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTurnCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetTurnCredentialsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTurnCredentialsRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetTurnCredentialsRequest) SetInstanceId(v string) *GetTurnCredentialsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTurnCredentialsRequest) SetUserId(v string) *GetTurnCredentialsRequest {
	s.UserId = &v
	return s
}

func (s *GetTurnCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
