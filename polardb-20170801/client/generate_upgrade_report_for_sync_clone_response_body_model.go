// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUpgradeReportForSyncCloneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateUpgradeReportForSyncCloneResponseBody
	GetRequestId() *string
	SetSourceDBClusterId(v string) *GenerateUpgradeReportForSyncCloneResponseBody
	GetSourceDBClusterId() *string
	SetTaskId(v int64) *GenerateUpgradeReportForSyncCloneResponseBody
	GetTaskId() *int64
}

type GenerateUpgradeReportForSyncCloneResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// pc-k2j96w169uhu868l8
	SourceDBClusterId *string `json:"SourceDBClusterId,omitempty" xml:"SourceDBClusterId,omitempty"`
	// example:
	//
	// 2312111
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GenerateUpgradeReportForSyncCloneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateUpgradeReportForSyncCloneResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUpgradeReportForSyncCloneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateUpgradeReportForSyncCloneResponseBody) GetSourceDBClusterId() *string {
	return s.SourceDBClusterId
}

func (s *GenerateUpgradeReportForSyncCloneResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GenerateUpgradeReportForSyncCloneResponseBody) SetRequestId(v string) *GenerateUpgradeReportForSyncCloneResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneResponseBody) SetSourceDBClusterId(v string) *GenerateUpgradeReportForSyncCloneResponseBody {
	s.SourceDBClusterId = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneResponseBody) SetTaskId(v int64) *GenerateUpgradeReportForSyncCloneResponseBody {
	s.TaskId = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneResponseBody) Validate() error {
	return dara.Validate(s)
}
