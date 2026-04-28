// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeltaGetLastCursorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v string) *DeltaGetLastCursorResponseBody
	GetCursor() *string
}

type DeltaGetLastCursorResponseBody struct {
	// The latest cursor of incremental information in the specified drive or synced folder.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
}

func (s DeltaGetLastCursorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeltaGetLastCursorResponseBody) GoString() string {
	return s.String()
}

func (s *DeltaGetLastCursorResponseBody) GetCursor() *string {
	return s.Cursor
}

func (s *DeltaGetLastCursorResponseBody) SetCursor(v string) *DeltaGetLastCursorResponseBody {
	s.Cursor = &v
	return s
}

func (s *DeltaGetLastCursorResponseBody) Validate() error {
	return dara.Validate(s)
}
