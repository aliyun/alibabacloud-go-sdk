// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *BatchStartApplicationsRequest
	GetAppIds() *string
	SetNamespaceId(v string) *BatchStartApplicationsRequest
	GetNamespaceId() *string
	SetVersion(v string) *BatchStartApplicationsRequest
	GetVersion() *string
}

type BatchStartApplicationsRequest struct {
	// The IDs of the applications that you want to start. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ebf491f0-c1a5-45e2-b2c4-710dbe2a****
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The application version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s BatchStartApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStartApplicationsRequest) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *BatchStartApplicationsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *BatchStartApplicationsRequest) GetVersion() *string {
	return s.Version
}

func (s *BatchStartApplicationsRequest) SetAppIds(v string) *BatchStartApplicationsRequest {
	s.AppIds = &v
	return s
}

func (s *BatchStartApplicationsRequest) SetNamespaceId(v string) *BatchStartApplicationsRequest {
	s.NamespaceId = &v
	return s
}

func (s *BatchStartApplicationsRequest) SetVersion(v string) *BatchStartApplicationsRequest {
	s.Version = &v
	return s
}

func (s *BatchStartApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
