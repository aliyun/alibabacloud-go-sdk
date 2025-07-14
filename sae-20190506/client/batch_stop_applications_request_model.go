// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *BatchStopApplicationsRequest
	GetAppIds() *string
	SetNamespaceId(v string) *BatchStopApplicationsRequest
	GetNamespaceId() *string
	SetVersion(v string) *BatchStopApplicationsRequest
	GetVersion() *string
}

type BatchStopApplicationsRequest struct {
	// The ID of the application that you want to stop.
	//
	// > If you want to stop multiple applications at the same time, separate the IDs with commas (,).
	//
	// example:
	//
	// ebf491f0-c1a5-45e2-b2c4-710dbe2a****
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// ebf491f0-c1a5-45e2-b2c4-710dbe2a\\*\\*\\*\\*,ebf491f0-c1a5-45e2-b2c4-71025e2a\\*\\*\\*\\*
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

func (s BatchStopApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStopApplicationsRequest) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *BatchStopApplicationsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *BatchStopApplicationsRequest) GetVersion() *string {
	return s.Version
}

func (s *BatchStopApplicationsRequest) SetAppIds(v string) *BatchStopApplicationsRequest {
	s.AppIds = &v
	return s
}

func (s *BatchStopApplicationsRequest) SetNamespaceId(v string) *BatchStopApplicationsRequest {
	s.NamespaceId = &v
	return s
}

func (s *BatchStopApplicationsRequest) SetVersion(v string) *BatchStopApplicationsRequest {
	s.Version = &v
	return s
}

func (s *BatchStopApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
