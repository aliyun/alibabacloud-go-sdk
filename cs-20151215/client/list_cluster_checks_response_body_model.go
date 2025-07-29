// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterChecksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChecks(v []*ListClusterChecksResponseBodyChecks) *ListClusterChecksResponseBody
	GetChecks() []*ListClusterChecksResponseBodyChecks
}

type ListClusterChecksResponseBody struct {
	// The list of check items.
	Checks []*ListClusterChecksResponseBodyChecks `json:"checks,omitempty" xml:"checks,omitempty" type:"Repeated"`
}

func (s ListClusterChecksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterChecksResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterChecksResponseBody) GetChecks() []*ListClusterChecksResponseBodyChecks {
	return s.Checks
}

func (s *ListClusterChecksResponseBody) SetChecks(v []*ListClusterChecksResponseBodyChecks) *ListClusterChecksResponseBody {
	s.Checks = v
	return s
}

func (s *ListClusterChecksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterChecksResponseBodyChecks struct {
	// The ID of the cluster check task.
	//
	// example:
	//
	// 1697100584236600453-ce0da5a1d627e4e9e9f96cae8ad07****-clustercheck-lboto
	CheckId *string `json:"check_id,omitempty" xml:"check_id,omitempty"`
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

func (s ListClusterChecksResponseBodyChecks) String() string {
	return dara.Prettify(s)
}

func (s ListClusterChecksResponseBodyChecks) GoString() string {
	return s.String()
}

func (s *ListClusterChecksResponseBodyChecks) GetCheckId() *string {
	return s.CheckId
}

func (s *ListClusterChecksResponseBodyChecks) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListClusterChecksResponseBodyChecks) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *ListClusterChecksResponseBodyChecks) GetMessage() *string {
	return s.Message
}

func (s *ListClusterChecksResponseBodyChecks) GetStatus() *string {
	return s.Status
}

func (s *ListClusterChecksResponseBodyChecks) GetType() *string {
	return s.Type
}

func (s *ListClusterChecksResponseBodyChecks) SetCheckId(v string) *ListClusterChecksResponseBodyChecks {
	s.CheckId = &v
	return s
}

func (s *ListClusterChecksResponseBodyChecks) SetCreatedAt(v string) *ListClusterChecksResponseBodyChecks {
	s.CreatedAt = &v
	return s
}

func (s *ListClusterChecksResponseBodyChecks) SetFinishedAt(v string) *ListClusterChecksResponseBodyChecks {
	s.FinishedAt = &v
	return s
}

func (s *ListClusterChecksResponseBodyChecks) SetMessage(v string) *ListClusterChecksResponseBodyChecks {
	s.Message = &v
	return s
}

func (s *ListClusterChecksResponseBodyChecks) SetStatus(v string) *ListClusterChecksResponseBodyChecks {
	s.Status = &v
	return s
}

func (s *ListClusterChecksResponseBodyChecks) SetType(v string) *ListClusterChecksResponseBodyChecks {
	s.Type = &v
	return s
}

func (s *ListClusterChecksResponseBodyChecks) Validate() error {
	return dara.Validate(s)
}
