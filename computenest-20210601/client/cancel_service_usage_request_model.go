// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelServiceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelServiceUsageRequest
	GetClientToken() *string
	SetNeedDelete(v bool) *CancelServiceUsageRequest
	GetNeedDelete() *bool
	SetServiceId(v string) *CancelServiceUsageRequest
	GetServiceId() *string
}

type CancelServiceUsageRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to delete the application.
	//
	// >  After you delete the application, you must re-enter the application information the next time you submit an application.
	//
	// example:
	//
	// true
	NeedDelete *bool `json:"NeedDelete,omitempty" xml:"NeedDelete,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-d6fc5f949a9246xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CancelServiceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *CancelServiceUsageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelServiceUsageRequest) GetNeedDelete() *bool {
	return s.NeedDelete
}

func (s *CancelServiceUsageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CancelServiceUsageRequest) SetClientToken(v string) *CancelServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelServiceUsageRequest) SetNeedDelete(v bool) *CancelServiceUsageRequest {
	s.NeedDelete = &v
	return s
}

func (s *CancelServiceUsageRequest) SetServiceId(v string) *CancelServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *CancelServiceUsageRequest) Validate() error {
	return dara.Validate(s)
}
