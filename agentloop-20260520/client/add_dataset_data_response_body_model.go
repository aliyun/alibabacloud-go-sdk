// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatasetDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedRows(v int32) *AddDatasetDataResponseBody
	GetAffectedRows() *int32
	SetRequestId(v string) *AddDatasetDataResponseBody
	GetRequestId() *string
}

type AddDatasetDataResponseBody struct {
	// The number of log rows scanned or processed.
	//
	// example:
	//
	// 100
	AffectedRows *int32 `json:"affectedRows,omitempty" xml:"affectedRows,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D0173835-9E0F-508F-8BFA-9F556E59C302
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddDatasetDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDataResponseBody) GoString() string {
	return s.String()
}

func (s *AddDatasetDataResponseBody) GetAffectedRows() *int32 {
	return s.AffectedRows
}

func (s *AddDatasetDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDatasetDataResponseBody) SetAffectedRows(v int32) *AddDatasetDataResponseBody {
	s.AffectedRows = &v
	return s
}

func (s *AddDatasetDataResponseBody) SetRequestId(v string) *AddDatasetDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDatasetDataResponseBody) Validate() error {
	return dara.Validate(s)
}
