// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetUserIdByOpenDingtalkIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *BatchGetUserIdByOpenDingtalkIdRequestTenantContext) *BatchGetUserIdByOpenDingtalkIdRequest
	GetTenantContext() *BatchGetUserIdByOpenDingtalkIdRequestTenantContext
	SetOpenDingtalkIds(v []*string) *BatchGetUserIdByOpenDingtalkIdRequest
	GetOpenDingtalkIds() []*string
}

type BatchGetUserIdByOpenDingtalkIdRequest struct {
	TenantContext *BatchGetUserIdByOpenDingtalkIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// ["DTOJdYJ2IQC4HuexhtjsS8Qxxxx","D8301AKf6lmZbXiilcB4P2MVxxxx"]
	OpenDingtalkIds []*string `json:"openDingtalkIds,omitempty" xml:"openDingtalkIds,omitempty" type:"Repeated"`
}

func (s BatchGetUserIdByOpenDingtalkIdRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetUserIdByOpenDingtalkIdRequest) GoString() string {
	return s.String()
}

func (s *BatchGetUserIdByOpenDingtalkIdRequest) GetTenantContext() *BatchGetUserIdByOpenDingtalkIdRequestTenantContext {
	return s.TenantContext
}

func (s *BatchGetUserIdByOpenDingtalkIdRequest) GetOpenDingtalkIds() []*string {
	return s.OpenDingtalkIds
}

func (s *BatchGetUserIdByOpenDingtalkIdRequest) SetTenantContext(v *BatchGetUserIdByOpenDingtalkIdRequestTenantContext) *BatchGetUserIdByOpenDingtalkIdRequest {
	s.TenantContext = v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdRequest) SetOpenDingtalkIds(v []*string) *BatchGetUserIdByOpenDingtalkIdRequest {
	s.OpenDingtalkIds = v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetUserIdByOpenDingtalkIdRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s BatchGetUserIdByOpenDingtalkIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s BatchGetUserIdByOpenDingtalkIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *BatchGetUserIdByOpenDingtalkIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchGetUserIdByOpenDingtalkIdRequestTenantContext) SetTenantId(v string) *BatchGetUserIdByOpenDingtalkIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
