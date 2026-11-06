// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterGroupChangeLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeLogs(v []*DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) *DescribeLakebaseMountParameterGroupChangeLogsResponseBody
	GetChangeLogs() []*DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs
	SetRequestId(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBody
	GetRequestId() *string
}

type DescribeLakebaseMountParameterGroupChangeLogsResponseBody struct {
	// The change records, sorted in descending order by CreateTime. An empty array is returned if no data is found.
	ChangeLogs []*DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs `json:"ChangeLogs,omitempty" xml:"ChangeLogs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupChangeLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupChangeLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBody) GetChangeLogs() []*DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs {
	return s.ChangeLogs
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBody) SetChangeLogs(v []*DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) *DescribeLakebaseMountParameterGroupChangeLogsResponseBody {
	s.ChangeLogs = v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBody) SetRequestId(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBody) Validate() error {
	if s.ChangeLogs != nil {
		for _, item := range s.ChangeLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs struct {
	// The change type. Valid values: create, modify, and delete.
	//
	// example:
	//
	// modify
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The creation time (UTC).
	//
	// example:
	//
	// 2026-08-07T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who performed the operation.
	//
	// example:
	//
	// 123456789
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Indicates whether the change has been applied. Valid values: 0 (no) and 1 (yes).
	//
	// example:
	//
	// 1
	IsApply *string `json:"IsApply,omitempty" xml:"IsApply,omitempty"`
	// The modification time (UTC).
	//
	// example:
	//
	// 2026-08-08T12:00:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The new value after the change. This parameter is empty for delete operations.
	//
	// example:
	//
	// 2
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// The old value before the change. This parameter is empty for create operations.
	//
	// example:
	//
	// 1
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// auto_increment_offset
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) GetChangeType() *string {
	return s.ChangeType
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) GetCreator() *string {
	return s.Creator
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) GetIsApply() *string {
	return s.IsApply
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) GetNewValue() *string {
	return s.NewValue
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) GetOldValue() *string {
	return s.OldValue
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) SetChangeType(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs {
	s.ChangeType = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) SetCreateTime(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs {
	s.CreateTime = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) SetCreator(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs {
	s.Creator = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) SetIsApply(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs {
	s.IsApply = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) SetModifyTime(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs {
	s.ModifyTime = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) SetNewValue(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs {
	s.NewValue = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) SetOldValue(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs {
	s.OldValue = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) SetParameterName(v string) *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs {
	s.ParameterName = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponseBodyChangeLogs) Validate() error {
	return dara.Validate(s)
}
