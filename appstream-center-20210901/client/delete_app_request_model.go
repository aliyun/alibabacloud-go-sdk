// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppRequest
	GetAppId() *string
	SetAppVersionId(v string) *DeleteAppRequest
	GetAppVersionId() *string
}

type DeleteAppRequest struct {
	// The ID of the application to be deleted. **This parameter is required**, and the application must belong to the current account. You must dissociate image references from the target application or version before deletion.
	//
	// example:
	//
	// app-bp1example****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application version to be deleted. If specified, only the adapters corresponding to this application version are deleted. If omitted, the entire application and its adapters are deleted. If the application has no other adapters after the version is deleted, the application record is also deleted.
	//
	// example:
	//
	// apv-bp1example****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppRequest) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *DeleteAppRequest) SetAppId(v string) *DeleteAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppRequest) SetAppVersionId(v string) *DeleteAppRequest {
	s.AppVersionId = &v
	return s
}

func (s *DeleteAppRequest) Validate() error {
	return dara.Validate(s)
}
