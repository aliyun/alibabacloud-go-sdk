// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteGrafanaResourceRequest
	GetClusterId() *string
	SetClusterName(v string) *DeleteGrafanaResourceRequest
	GetClusterName() *string
	SetUserId(v string) *DeleteGrafanaResourceRequest
	GetUserId() *string
}

type DeleteGrafanaResourceRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteGrafanaResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteGrafanaResourceRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DeleteGrafanaResourceRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteGrafanaResourceRequest) SetClusterId(v string) *DeleteGrafanaResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) SetClusterName(v string) *DeleteGrafanaResourceRequest {
	s.ClusterName = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) SetUserId(v string) *DeleteGrafanaResourceRequest {
	s.UserId = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) Validate() error {
	return dara.Validate(s)
}
