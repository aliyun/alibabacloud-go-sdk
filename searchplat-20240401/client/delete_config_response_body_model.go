// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConfigResponseBody
	GetRequestId() *string
	SetResult(v *DeleteConfigResponseBodyResult) *DeleteConfigResponseBody
	GetResult() *DeleteConfigResponseBodyResult
}

type DeleteConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CC93E65-6734-5060-BEF7-0EB0A4862BCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *DeleteConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DeleteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConfigResponseBody) GetResult() *DeleteConfigResponseBodyResult {
	return s.Result
}

func (s *DeleteConfigResponseBody) SetRequestId(v string) *DeleteConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigResponseBody) SetResult(v *DeleteConfigResponseBodyResult) *DeleteConfigResponseBody {
	s.Result = v
	return s
}

func (s *DeleteConfigResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteConfigResponseBodyResult struct {
	// The configuration type.
	//
	// - prompt
	//
	// - lark
	//
	// example:
	//
	// prompt
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// Indicates whether the configuration is deleted.
	//
	// example:
	//
	// true
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// a1b2c3
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1200827
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteConfigResponseBodyResult) GetConfigType() *string {
	return s.ConfigType
}

func (s *DeleteConfigResponseBodyResult) GetDeleted() *bool {
	return s.Deleted
}

func (s *DeleteConfigResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DeleteConfigResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteConfigResponseBodyResult) SetConfigType(v string) *DeleteConfigResponseBodyResult {
	s.ConfigType = &v
	return s
}

func (s *DeleteConfigResponseBodyResult) SetDeleted(v bool) *DeleteConfigResponseBodyResult {
	s.Deleted = &v
	return s
}

func (s *DeleteConfigResponseBodyResult) SetId(v string) *DeleteConfigResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DeleteConfigResponseBodyResult) SetWorkspaceId(v string) *DeleteConfigResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
