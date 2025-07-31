// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferClusterBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlreadyDone(v string) *TransferClusterBackupResponseBody
	GetAlreadyDone() *string
	SetRequestId(v string) *TransferClusterBackupResponseBody
	GetRequestId() *string
}

type TransferClusterBackupResponseBody struct {
	// Indicates whether the instance is switched to the cluster backup mode. If the value of this parameter is **1**, the instance is switched to the cluster backup mode.
	//
	// example:
	//
	// 1
	AlreadyDone *string `json:"AlreadyDone,omitempty" xml:"AlreadyDone,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C4A2494-85C4-45C5-93CF-548DB3375193
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferClusterBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferClusterBackupResponseBody) GoString() string {
	return s.String()
}

func (s *TransferClusterBackupResponseBody) GetAlreadyDone() *string {
	return s.AlreadyDone
}

func (s *TransferClusterBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferClusterBackupResponseBody) SetAlreadyDone(v string) *TransferClusterBackupResponseBody {
	s.AlreadyDone = &v
	return s
}

func (s *TransferClusterBackupResponseBody) SetRequestId(v string) *TransferClusterBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferClusterBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
