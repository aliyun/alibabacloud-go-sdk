// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceMigrationAvailableNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReplaceMigrationAvailableNumbersRequest
	GetInstanceId() *string
	SetOperatorName(v string) *ReplaceMigrationAvailableNumbersRequest
	GetOperatorName() *string
	SetOperatorUid(v int64) *ReplaceMigrationAvailableNumbersRequest
	GetOperatorUid() *int64
	SetV1Numbers(v string) *ReplaceMigrationAvailableNumbersRequest
	GetV1Numbers() *string
	SetV2Numbers(v string) *ReplaceMigrationAvailableNumbersRequest
	GetV2Numbers() *string
}

type ReplaceMigrationAvailableNumbersRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	OperatorUid  *int64  `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	V1Numbers    *string `json:"V1Numbers,omitempty" xml:"V1Numbers,omitempty"`
	V2Numbers    *string `json:"V2Numbers,omitempty" xml:"V2Numbers,omitempty"`
}

func (s ReplaceMigrationAvailableNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceMigrationAvailableNumbersRequest) GoString() string {
	return s.String()
}

func (s *ReplaceMigrationAvailableNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReplaceMigrationAvailableNumbersRequest) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ReplaceMigrationAvailableNumbersRequest) GetOperatorUid() *int64 {
	return s.OperatorUid
}

func (s *ReplaceMigrationAvailableNumbersRequest) GetV1Numbers() *string {
	return s.V1Numbers
}

func (s *ReplaceMigrationAvailableNumbersRequest) GetV2Numbers() *string {
	return s.V2Numbers
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetInstanceId(v string) *ReplaceMigrationAvailableNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetOperatorName(v string) *ReplaceMigrationAvailableNumbersRequest {
	s.OperatorName = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetOperatorUid(v int64) *ReplaceMigrationAvailableNumbersRequest {
	s.OperatorUid = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetV1Numbers(v string) *ReplaceMigrationAvailableNumbersRequest {
	s.V1Numbers = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetV2Numbers(v string) *ReplaceMigrationAvailableNumbersRequest {
	s.V2Numbers = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersRequest) Validate() error {
	return dara.Validate(s)
}
