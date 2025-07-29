// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v string) *GetClusterCheckResponseBody
	GetCheckId() *string
	SetCheckItems(v map[string][]map[string]interface{}) *GetClusterCheckResponseBody
	GetCheckItems() map[string][]map[string]interface{}
	SetCreatedAt(v string) *GetClusterCheckResponseBody
	GetCreatedAt() *string
	SetFinishedAt(v string) *GetClusterCheckResponseBody
	GetFinishedAt() *string
	SetMessage(v string) *GetClusterCheckResponseBody
	GetMessage() *string
	SetStatus(v string) *GetClusterCheckResponseBody
	GetStatus() *string
	SetType(v string) *GetClusterCheckResponseBody
	GetType() *string
}

type GetClusterCheckResponseBody struct {
	// The ID of the cluster check task.
	//
	// example:
	//
	// 1697100584236600453-ce0da5a1d627e4e9e9f96cae8ad07****-clustercheck-lboto
	CheckId *string `json:"check_id,omitempty" xml:"check_id,omitempty"`
	// A list of check items.
	CheckItems map[string][]map[string]interface{} `json:"check_items,omitempty" xml:"check_items,omitempty"`
	// The time when the cluster check task was created.
	//
	// example:
	//
	// 2023-10-16T08:31:20.292030178Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The time when the cluster check task was completed.
	//
	// example:
	//
	// 2023-10-16T08:35:20.292030178Z
	FinishedAt *string `json:"finished_at,omitempty" xml:"finished_at,omitempty"`
	// The message that indicates the status of the cluster check task.
	//
	// example:
	//
	// task succeed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The status of the cluster check.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The check method.
	//
	// example:
	//
	// ClusterUpgrade
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetClusterCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterCheckResponseBody) GetCheckId() *string {
	return s.CheckId
}

func (s *GetClusterCheckResponseBody) GetCheckItems() map[string][]map[string]interface{} {
	return s.CheckItems
}

func (s *GetClusterCheckResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetClusterCheckResponseBody) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetClusterCheckResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetClusterCheckResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetClusterCheckResponseBody) GetType() *string {
	return s.Type
}

func (s *GetClusterCheckResponseBody) SetCheckId(v string) *GetClusterCheckResponseBody {
	s.CheckId = &v
	return s
}

func (s *GetClusterCheckResponseBody) SetCheckItems(v map[string][]map[string]interface{}) *GetClusterCheckResponseBody {
	s.CheckItems = v
	return s
}

func (s *GetClusterCheckResponseBody) SetCreatedAt(v string) *GetClusterCheckResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetClusterCheckResponseBody) SetFinishedAt(v string) *GetClusterCheckResponseBody {
	s.FinishedAt = &v
	return s
}

func (s *GetClusterCheckResponseBody) SetMessage(v string) *GetClusterCheckResponseBody {
	s.Message = &v
	return s
}

func (s *GetClusterCheckResponseBody) SetStatus(v string) *GetClusterCheckResponseBody {
	s.Status = &v
	return s
}

func (s *GetClusterCheckResponseBody) SetType(v string) *GetClusterCheckResponseBody {
	s.Type = &v
	return s
}

func (s *GetClusterCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
