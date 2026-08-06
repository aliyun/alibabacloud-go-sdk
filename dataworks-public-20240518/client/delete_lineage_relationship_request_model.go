// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLineageRelationshipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteLineageRelationshipRequest
	GetId() *string
}

type DeleteLineageRelationshipRequest struct {
	// The ID of the lineage relationship. You can obtain this ID from the response of the ListLineageRelationships operation. The ID is in the format of `${accountId}:${srcEntityType}.${srcEntityId}:${dstEntityType}.${dstEntityId}:${taskType}.${taskId}`, where accountId is the Alibaba Cloud account ID, srcEntityType and srcEntityId are the source entity type and source entity ID, dstEntityType and dstEntityId are the destination entity type and destination entity ID, and taskType and taskId are the lineage task type and task ID. Example: `1245491995595649:custom-report.report_test_001:custom-table.table_test_001:custom-lineage-task.test_task_001`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4as3dasf654a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteLineageRelationshipRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLineageRelationshipRequest) GoString() string {
	return s.String()
}

func (s *DeleteLineageRelationshipRequest) GetId() *string {
	return s.Id
}

func (s *DeleteLineageRelationshipRequest) SetId(v string) *DeleteLineageRelationshipRequest {
	s.Id = &v
	return s
}

func (s *DeleteLineageRelationshipRequest) Validate() error {
	return dara.Validate(s)
}
