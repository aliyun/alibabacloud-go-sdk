// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordEnableStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionRecordEnableStatusRequest
	GetClientToken() *string
	SetEnableStatus(v string) *UpdateRecursionRecordEnableStatusRequest
	GetEnableStatus() *string
	SetRecordId(v string) *UpdateRecursionRecordEnableStatusRequest
	GetRecordId() *string
}

type UpdateRecursionRecordEnableStatusRequest struct {
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// example:
	//
	// 1781234321
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s UpdateRecursionRecordEnableStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordEnableStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionRecordEnableStatusRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *UpdateRecursionRecordEnableStatusRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateRecursionRecordEnableStatusRequest) SetClientToken(v string) *UpdateRecursionRecordEnableStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionRecordEnableStatusRequest) SetEnableStatus(v string) *UpdateRecursionRecordEnableStatusRequest {
	s.EnableStatus = &v
	return s
}

func (s *UpdateRecursionRecordEnableStatusRequest) SetRecordId(v string) *UpdateRecursionRecordEnableStatusRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecursionRecordEnableStatusRequest) Validate() error {
	return dara.Validate(s)
}
