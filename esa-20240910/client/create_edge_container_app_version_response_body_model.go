// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEdgeContainerAppVersionResponseBody
	GetRequestId() *string
	SetVersionId(v string) *CreateEdgeContainerAppVersionResponseBody
	GetVersionId() *string
}

type CreateEdgeContainerAppVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created version.
	//
	// example:
	//
	// ver-87962637161651****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateEdgeContainerAppVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEdgeContainerAppVersionResponseBody) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateEdgeContainerAppVersionResponseBody) SetRequestId(v string) *CreateEdgeContainerAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEdgeContainerAppVersionResponseBody) SetVersionId(v string) *CreateEdgeContainerAppVersionResponseBody {
	s.VersionId = &v
	return s
}

func (s *CreateEdgeContainerAppVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
