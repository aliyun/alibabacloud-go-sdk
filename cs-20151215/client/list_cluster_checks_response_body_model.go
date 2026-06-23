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
	// The list of checks.
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
	if s.Checks != nil {
		for _, item := range s.Checks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterChecksResponseBodyChecks struct {
	// The check ID.
	//
	// example:
	//
	// 1697100584236600453-ce0da5a1d627e4e9e9f96cae8ad07****-clustercheck-lboto
	CheckId *string `json:"check_id,omitempty" xml:"check_id,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-04-11T02:56:02.565982623Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The completion time.
	//
	// example:
	//
	// 2025-04-11T02:56:18.881054031Z
	FinishedAt *string `json:"finished_at,omitempty" xml:"finished_at,omitempty"`
	// The check status message.
	//
	// example:
	//
	// task succeed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The check status.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The check type.
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
