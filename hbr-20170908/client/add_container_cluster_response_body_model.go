// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AddContainerClusterResponseBody
	GetClusterId() *string
	SetCode(v string) *AddContainerClusterResponseBody
	GetCode() *string
	SetMessage(v string) *AddContainerClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddContainerClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddContainerClusterResponseBody
	GetSuccess() *bool
	SetToken(v string) *AddContainerClusterResponseBody
	GetToken() *string
}

type AddContainerClusterResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// cc-00049slr9iuvvv6pp134
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the request is successful, a value of successful is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1FCBC078-FFCB-542A-8555-566477679720
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The token that is used to register the Hybrid Backup Recovery (HBR) client in the cluster.
	//
	// example:
	//
	// eyJhY2NvdW*****VnZpgXQC5A==
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s AddContainerClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddContainerClusterResponseBody) GoString() string {
	return s.String()
}

func (s *AddContainerClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddContainerClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddContainerClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddContainerClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddContainerClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddContainerClusterResponseBody) GetToken() *string {
	return s.Token
}

func (s *AddContainerClusterResponseBody) SetClusterId(v string) *AddContainerClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetCode(v string) *AddContainerClusterResponseBody {
	s.Code = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetMessage(v string) *AddContainerClusterResponseBody {
	s.Message = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetRequestId(v string) *AddContainerClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetSuccess(v bool) *AddContainerClusterResponseBody {
	s.Success = &v
	return s
}

func (s *AddContainerClusterResponseBody) SetToken(v string) *AddContainerClusterResponseBody {
	s.Token = &v
	return s
}

func (s *AddContainerClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
