// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaTableIntroWikiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMetaTableIntroWikiResponseBody
	GetRequestId() *string
	SetUpdateResult(v bool) *UpdateMetaTableIntroWikiResponseBody
	GetUpdateResult() *bool
}

type UpdateMetaTableIntroWikiResponseBody struct {
	// The unique ID of the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The update result of the table.
	//
	// example:
	//
	// true
	UpdateResult *bool `json:"UpdateResult,omitempty" xml:"UpdateResult,omitempty"`
}

func (s UpdateMetaTableIntroWikiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaTableIntroWikiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaTableIntroWikiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetaTableIntroWikiResponseBody) GetUpdateResult() *bool {
	return s.UpdateResult
}

func (s *UpdateMetaTableIntroWikiResponseBody) SetRequestId(v string) *UpdateMetaTableIntroWikiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetaTableIntroWikiResponseBody) SetUpdateResult(v bool) *UpdateMetaTableIntroWikiResponseBody {
	s.UpdateResult = &v
	return s
}

func (s *UpdateMetaTableIntroWikiResponseBody) Validate() error {
	return dara.Validate(s)
}
