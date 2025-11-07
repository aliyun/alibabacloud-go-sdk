// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QuerySceneConfigsResponseBody
	GetRequestId() *string
	SetSceneConfigs(v []*QuerySceneConfigsResponseBodySceneConfigs) *QuerySceneConfigsResponseBody
	GetSceneConfigs() []*QuerySceneConfigsResponseBodySceneConfigs
}

type QuerySceneConfigsResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Willingness configuration list.
	SceneConfigs []*QuerySceneConfigsResponseBodySceneConfigs `json:"sceneConfigs,omitempty" xml:"sceneConfigs,omitempty" type:"Repeated"`
}

func (s QuerySceneConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySceneConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySceneConfigsResponseBody) GetSceneConfigs() []*QuerySceneConfigsResponseBodySceneConfigs {
	return s.SceneConfigs
}

func (s *QuerySceneConfigsResponseBody) SetRequestId(v string) *QuerySceneConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySceneConfigsResponseBody) SetSceneConfigs(v []*QuerySceneConfigsResponseBodySceneConfigs) *QuerySceneConfigsResponseBody {
	s.SceneConfigs = v
	return s
}

func (s *QuerySceneConfigsResponseBody) Validate() error {
	if s.SceneConfigs != nil {
		for _, item := range s.SceneConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySceneConfigsResponseBodySceneConfigs struct {
	// Specific configuration content, in JSON string format.
	//
	// example:
	//
	// {\\"faceCompareMode\\":\\"CUSTOM\\",\\"certConfigs\\":[{\\"index\\":0,\\"openVoiceCompare\\":false,\\"openCustomizedContent\\":true,\\"model\\":\\"QA\\"}],\\"screenEvidence\\":true}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1760782820000
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1760782820000
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Configuration ID.
	//
	// example:
	//
	// 607
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 1000009045
	SceneId *int64 `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	// Configuration type.
	//
	// example:
	//
	// VOLUNTARY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Scene configuration version number.
	//
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QuerySceneConfigsResponseBodySceneConfigs) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneConfigsResponseBodySceneConfigs) GoString() string {
	return s.String()
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) GetConfig() *string {
	return s.Config
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) GetId() *int64 {
	return s.Id
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) GetSceneId() *int64 {
	return s.SceneId
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) GetType() *string {
	return s.Type
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) GetVersion() *int64 {
	return s.Version
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) SetConfig(v string) *QuerySceneConfigsResponseBodySceneConfigs {
	s.Config = &v
	return s
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) SetGmtCreate(v string) *QuerySceneConfigsResponseBodySceneConfigs {
	s.GmtCreate = &v
	return s
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) SetGmtModified(v string) *QuerySceneConfigsResponseBodySceneConfigs {
	s.GmtModified = &v
	return s
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) SetId(v int64) *QuerySceneConfigsResponseBodySceneConfigs {
	s.Id = &v
	return s
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) SetSceneId(v int64) *QuerySceneConfigsResponseBodySceneConfigs {
	s.SceneId = &v
	return s
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) SetType(v string) *QuerySceneConfigsResponseBodySceneConfigs {
	s.Type = &v
	return s
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) SetVersion(v int64) *QuerySceneConfigsResponseBodySceneConfigs {
	s.Version = &v
	return s
}

func (s *QuerySceneConfigsResponseBodySceneConfigs) Validate() error {
	return dara.Validate(s)
}
