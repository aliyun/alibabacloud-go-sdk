// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGraphInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessProfile(v string) *UpdateGraphInfoResponseBody
	GetBusinessProfile() *string
	SetCode(v string) *UpdateGraphInfoResponseBody
	GetCode() *string
	SetDisplayName(v string) *UpdateGraphInfoResponseBody
	GetDisplayName() *string
	SetGraphName(v string) *UpdateGraphInfoResponseBody
	GetGraphName() *string
	SetMessage(v string) *UpdateGraphInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGraphInfoResponseBody
	GetRequestId() *string
	SetUpdated(v bool) *UpdateGraphInfoResponseBody
	GetUpdated() *bool
}

type UpdateGraphInfoResponseBody struct {
	// The business description of the knowledge graph. If not configured, the value is an empty string.
	//
	// example:
	//
	// Customer domain knowledge graph
	BusinessProfile *string `json:"businessProfile,omitempty" xml:"businessProfile,omitempty"`
	// SUCCESS indicates success. In failure cases, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The display name of the knowledge graph.
	//
	// example:
	//
	// CRM Graph
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the knowledge graph.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The updated items.
	//
	// example:
	//
	// true
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateGraphInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGraphInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGraphInfoResponseBody) GetBusinessProfile() *string {
	return s.BusinessProfile
}

func (s *UpdateGraphInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGraphInfoResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateGraphInfoResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *UpdateGraphInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGraphInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGraphInfoResponseBody) GetUpdated() *bool {
	return s.Updated
}

func (s *UpdateGraphInfoResponseBody) SetBusinessProfile(v string) *UpdateGraphInfoResponseBody {
	s.BusinessProfile = &v
	return s
}

func (s *UpdateGraphInfoResponseBody) SetCode(v string) *UpdateGraphInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGraphInfoResponseBody) SetDisplayName(v string) *UpdateGraphInfoResponseBody {
	s.DisplayName = &v
	return s
}

func (s *UpdateGraphInfoResponseBody) SetGraphName(v string) *UpdateGraphInfoResponseBody {
	s.GraphName = &v
	return s
}

func (s *UpdateGraphInfoResponseBody) SetMessage(v string) *UpdateGraphInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGraphInfoResponseBody) SetRequestId(v string) *UpdateGraphInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGraphInfoResponseBody) SetUpdated(v bool) *UpdateGraphInfoResponseBody {
	s.Updated = &v
	return s
}

func (s *UpdateGraphInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
