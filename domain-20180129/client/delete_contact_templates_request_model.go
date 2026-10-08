// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegistrantProfileIds(v string) *DeleteContactTemplatesRequest
	GetRegistrantProfileIds() *string
	SetUserClientIp(v string) *DeleteContactTemplatesRequest
	GetUserClientIp() *string
}

type DeleteContactTemplatesRequest struct {
	// The IDs of the contact templates to delete. Separate multiple values with commas (,).
	//
	// The system automatically generates an ID upon successful creation of a contact template. You can invoke the [QueryRegistrantProfiles](https://help.aliyun.com/document_detail/67701.html) API to query the template IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123,45,67
	RegistrantProfileIds *string `json:"RegistrantProfileIds,omitempty" xml:"RegistrantProfileIds,omitempty"`
	// User IP address. You can set this parameter to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteContactTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactTemplatesRequest) GetRegistrantProfileIds() *string {
	return s.RegistrantProfileIds
}

func (s *DeleteContactTemplatesRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DeleteContactTemplatesRequest) SetRegistrantProfileIds(v string) *DeleteContactTemplatesRequest {
	s.RegistrantProfileIds = &v
	return s
}

func (s *DeleteContactTemplatesRequest) SetUserClientIp(v string) *DeleteContactTemplatesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteContactTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
