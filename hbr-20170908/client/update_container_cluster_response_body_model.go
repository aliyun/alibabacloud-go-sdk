// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContainerClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateContainerClusterResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateContainerClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContainerClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateContainerClusterResponseBody
	GetSuccess() *bool
	SetToken(v string) *UpdateContainerClusterResponseBody
	GetToken() *string
	SetTokenUpdated(v bool) *UpdateContainerClusterResponseBody
	GetTokenUpdated() *bool
}

type UpdateContainerClusterResponseBody struct {
	// Return code, 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return information.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates if the request was successful.
	//
	// - true: Success
	//
	// - false: Failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Cluster token, used for registering HBR clients within the cluster.
	//
	// example:
	//
	// eyJhY2NvdW*****VnZpgXQC5A==
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// Indicates whether the cluster token has been updated.
	//
	// example:
	//
	// false
	TokenUpdated *bool `json:"TokenUpdated,omitempty" xml:"TokenUpdated,omitempty"`
}

func (s UpdateContainerClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContainerClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateContainerClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContainerClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContainerClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateContainerClusterResponseBody) GetToken() *string {
	return s.Token
}

func (s *UpdateContainerClusterResponseBody) GetTokenUpdated() *bool {
	return s.TokenUpdated
}

func (s *UpdateContainerClusterResponseBody) SetCode(v string) *UpdateContainerClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContainerClusterResponseBody) SetMessage(v string) *UpdateContainerClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContainerClusterResponseBody) SetRequestId(v string) *UpdateContainerClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContainerClusterResponseBody) SetSuccess(v bool) *UpdateContainerClusterResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateContainerClusterResponseBody) SetToken(v string) *UpdateContainerClusterResponseBody {
	s.Token = &v
	return s
}

func (s *UpdateContainerClusterResponseBody) SetTokenUpdated(v bool) *UpdateContainerClusterResponseBody {
	s.TokenUpdated = &v
	return s
}

func (s *UpdateContainerClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
