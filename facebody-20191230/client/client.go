// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	number "github.com/alibabacloud-go/darabonba-number/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type AddFaceRequest struct {
	DbName                                *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId                              *string  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ExtraData                             *string  `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	ImageUrl                              *string  `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	QualityScoreThreshold                 *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
	SimilarityScoreThresholdBetweenEntity *float32 `json:"SimilarityScoreThresholdBetweenEntity,omitempty" xml:"SimilarityScoreThresholdBetweenEntity,omitempty"`
	SimilarityScoreThresholdInEntity      *float32 `json:"SimilarityScoreThresholdInEntity,omitempty" xml:"SimilarityScoreThresholdInEntity,omitempty"`
}

func (s AddFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceRequest) SetDbName(v string) *AddFaceRequest {
	s.DbName = &v
	return s
}

func (s *AddFaceRequest) SetEntityId(v string) *AddFaceRequest {
	s.EntityId = &v
	return s
}

func (s *AddFaceRequest) SetExtraData(v string) *AddFaceRequest {
	s.ExtraData = &v
	return s
}

func (s *AddFaceRequest) SetImageUrl(v string) *AddFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *AddFaceRequest) SetQualityScoreThreshold(v float32) *AddFaceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *AddFaceRequest) SetSimilarityScoreThresholdBetweenEntity(v float32) *AddFaceRequest {
	s.SimilarityScoreThresholdBetweenEntity = &v
	return s
}

func (s *AddFaceRequest) SetSimilarityScoreThresholdInEntity(v float32) *AddFaceRequest {
	s.SimilarityScoreThresholdInEntity = &v
	return s
}

type AddFaceAdvanceRequest struct {
	DbName                                *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId                              *string   `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ExtraData                             *string   `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	ImageUrlObject                        io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	QualityScoreThreshold                 *float32  `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
	SimilarityScoreThresholdBetweenEntity *float32  `json:"SimilarityScoreThresholdBetweenEntity,omitempty" xml:"SimilarityScoreThresholdBetweenEntity,omitempty"`
	SimilarityScoreThresholdInEntity      *float32  `json:"SimilarityScoreThresholdInEntity,omitempty" xml:"SimilarityScoreThresholdInEntity,omitempty"`
}

func (s AddFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceAdvanceRequest) SetDbName(v string) *AddFaceAdvanceRequest {
	s.DbName = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetEntityId(v string) *AddFaceAdvanceRequest {
	s.EntityId = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetExtraData(v string) *AddFaceAdvanceRequest {
	s.ExtraData = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetImageUrlObject(v io.Reader) *AddFaceAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *AddFaceAdvanceRequest) SetQualityScoreThreshold(v float32) *AddFaceAdvanceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetSimilarityScoreThresholdBetweenEntity(v float32) *AddFaceAdvanceRequest {
	s.SimilarityScoreThresholdBetweenEntity = &v
	return s
}

func (s *AddFaceAdvanceRequest) SetSimilarityScoreThresholdInEntity(v float32) *AddFaceAdvanceRequest {
	s.SimilarityScoreThresholdInEntity = &v
	return s
}

type AddFaceResponseBody struct {
	Data      *AddFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceResponseBody) SetData(v *AddFaceResponseBodyData) *AddFaceResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceResponseBody) SetRequestId(v string) *AddFaceResponseBody {
	s.RequestId = &v
	return s
}

type AddFaceResponseBodyData struct {
	FaceId        *string  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	QualitieScore *float32 `json:"QualitieScore,omitempty" xml:"QualitieScore,omitempty"`
}

func (s AddFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceResponseBodyData) SetFaceId(v string) *AddFaceResponseBodyData {
	s.FaceId = &v
	return s
}

func (s *AddFaceResponseBodyData) SetQualitieScore(v float32) *AddFaceResponseBodyData {
	s.QualitieScore = &v
	return s
}

type AddFaceResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceResponse) GoString() string {
	return s.String()
}

func (s *AddFaceResponse) SetHeaders(v map[string]*string) *AddFaceResponse {
	s.Headers = v
	return s
}

func (s *AddFaceResponse) SetStatusCode(v int32) *AddFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceResponse) SetBody(v *AddFaceResponseBody) *AddFaceResponse {
	s.Body = v
	return s
}

type AddFaceEntityRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Labels   *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s AddFaceEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *AddFaceEntityRequest) SetDbName(v string) *AddFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *AddFaceEntityRequest) SetEntityId(v string) *AddFaceEntityRequest {
	s.EntityId = &v
	return s
}

func (s *AddFaceEntityRequest) SetLabels(v string) *AddFaceEntityRequest {
	s.Labels = &v
	return s
}

type AddFaceEntityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFaceEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceEntityResponseBody) SetRequestId(v string) *AddFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

type AddFaceEntityResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFaceEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *AddFaceEntityResponse) SetHeaders(v map[string]*string) *AddFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *AddFaceEntityResponse) SetStatusCode(v int32) *AddFaceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceEntityResponse) SetBody(v *AddFaceEntityResponseBody) *AddFaceEntityResponse {
	s.Body = v
	return s
}

type AddFaceImageTemplateRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AddFaceImageTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateRequest) SetImageURL(v string) *AddFaceImageTemplateRequest {
	s.ImageURL = &v
	return s
}

type AddFaceImageTemplateAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AddFaceImageTemplateAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateAdvanceRequest) SetImageURLObject(v io.Reader) *AddFaceImageTemplateAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type AddFaceImageTemplateResponseBody struct {
	Data      *AddFaceImageTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFaceImageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponseBody) SetData(v *AddFaceImageTemplateResponseBodyData) *AddFaceImageTemplateResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceImageTemplateResponseBody) SetRequestId(v string) *AddFaceImageTemplateResponseBody {
	s.RequestId = &v
	return s
}

type AddFaceImageTemplateResponseBodyData struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddFaceImageTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponseBodyData) SetTemplateId(v string) *AddFaceImageTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

type AddFaceImageTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddFaceImageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFaceImageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFaceImageTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponse) SetHeaders(v map[string]*string) *AddFaceImageTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddFaceImageTemplateResponse) SetStatusCode(v int32) *AddFaceImageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceImageTemplateResponse) SetBody(v *AddFaceImageTemplateResponseBody) *AddFaceImageTemplateResponse {
	s.Body = v
	return s
}

type BatchAddFacesRequest struct {
	DbName                                *string                      `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId                              *string                      `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Faces                                 []*BatchAddFacesRequestFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	QualityScoreThreshold                 *float32                     `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
	SimilarityScoreThresholdBetweenEntity *float32                     `json:"SimilarityScoreThresholdBetweenEntity,omitempty" xml:"SimilarityScoreThresholdBetweenEntity,omitempty"`
	SimilarityScoreThresholdInEntity      *float32                     `json:"SimilarityScoreThresholdInEntity,omitempty" xml:"SimilarityScoreThresholdInEntity,omitempty"`
}

func (s BatchAddFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFacesRequest) GoString() string {
	return s.String()
}

func (s *BatchAddFacesRequest) SetDbName(v string) *BatchAddFacesRequest {
	s.DbName = &v
	return s
}

func (s *BatchAddFacesRequest) SetEntityId(v string) *BatchAddFacesRequest {
	s.EntityId = &v
	return s
}

func (s *BatchAddFacesRequest) SetFaces(v []*BatchAddFacesRequestFaces) *BatchAddFacesRequest {
	s.Faces = v
	return s
}

func (s *BatchAddFacesRequest) SetQualityScoreThreshold(v float32) *BatchAddFacesRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *BatchAddFacesRequest) SetSimilarityScoreThresholdBetweenEntity(v float32) *BatchAddFacesRequest {
	s.SimilarityScoreThresholdBetweenEntity = &v
	return s
}

func (s *BatchAddFacesRequest) SetSimilarityScoreThresholdInEntity(v float32) *BatchAddFacesRequest {
	s.SimilarityScoreThresholdInEntity = &v
	return s
}

type BatchAddFacesRequestFaces struct {
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BatchAddFacesRequestFaces) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFacesRequestFaces) GoString() string {
	return s.String()
}

func (s *BatchAddFacesRequestFaces) SetExtraData(v string) *BatchAddFacesRequestFaces {
	s.ExtraData = &v
	return s
}

func (s *BatchAddFacesRequestFaces) SetImageURL(v string) *BatchAddFacesRequestFaces {
	s.ImageURL = &v
	return s
}

type BatchAddFacesShrinkRequest struct {
	DbName                                *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId                              *string  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	FacesShrink                           *string  `json:"Faces,omitempty" xml:"Faces,omitempty"`
	QualityScoreThreshold                 *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
	SimilarityScoreThresholdBetweenEntity *float32 `json:"SimilarityScoreThresholdBetweenEntity,omitempty" xml:"SimilarityScoreThresholdBetweenEntity,omitempty"`
	SimilarityScoreThresholdInEntity      *float32 `json:"SimilarityScoreThresholdInEntity,omitempty" xml:"SimilarityScoreThresholdInEntity,omitempty"`
}

func (s BatchAddFacesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchAddFacesShrinkRequest) SetDbName(v string) *BatchAddFacesShrinkRequest {
	s.DbName = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetEntityId(v string) *BatchAddFacesShrinkRequest {
	s.EntityId = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetFacesShrink(v string) *BatchAddFacesShrinkRequest {
	s.FacesShrink = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetQualityScoreThreshold(v float32) *BatchAddFacesShrinkRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetSimilarityScoreThresholdBetweenEntity(v float32) *BatchAddFacesShrinkRequest {
	s.SimilarityScoreThresholdBetweenEntity = &v
	return s
}

func (s *BatchAddFacesShrinkRequest) SetSimilarityScoreThresholdInEntity(v float32) *BatchAddFacesShrinkRequest {
	s.SimilarityScoreThresholdInEntity = &v
	return s
}

type BatchAddFacesResponseBody struct {
	Data      *BatchAddFacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchAddFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFacesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponseBody) SetData(v *BatchAddFacesResponseBodyData) *BatchAddFacesResponseBody {
	s.Data = v
	return s
}

func (s *BatchAddFacesResponseBody) SetRequestId(v string) *BatchAddFacesResponseBody {
	s.RequestId = &v
	return s
}

type BatchAddFacesResponseBodyData struct {
	FailedFaces   []*BatchAddFacesResponseBodyDataFailedFaces   `json:"FailedFaces,omitempty" xml:"FailedFaces,omitempty" type:"Repeated"`
	InsertedFaces []*BatchAddFacesResponseBodyDataInsertedFaces `json:"InsertedFaces,omitempty" xml:"InsertedFaces,omitempty" type:"Repeated"`
}

func (s BatchAddFacesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponseBodyData) SetFailedFaces(v []*BatchAddFacesResponseBodyDataFailedFaces) *BatchAddFacesResponseBodyData {
	s.FailedFaces = v
	return s
}

func (s *BatchAddFacesResponseBodyData) SetInsertedFaces(v []*BatchAddFacesResponseBodyDataInsertedFaces) *BatchAddFacesResponseBodyData {
	s.InsertedFaces = v
	return s
}

type BatchAddFacesResponseBodyDataFailedFaces struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s BatchAddFacesResponseBodyDataFailedFaces) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFacesResponseBodyDataFailedFaces) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) SetCode(v string) *BatchAddFacesResponseBodyDataFailedFaces {
	s.Code = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) SetImageURL(v string) *BatchAddFacesResponseBodyDataFailedFaces {
	s.ImageURL = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataFailedFaces) SetMessage(v string) *BatchAddFacesResponseBodyDataFailedFaces {
	s.Message = &v
	return s
}

type BatchAddFacesResponseBodyDataInsertedFaces struct {
	FaceId        *string  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	ImageURL      *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	QualitieScore *float32 `json:"QualitieScore,omitempty" xml:"QualitieScore,omitempty"`
}

func (s BatchAddFacesResponseBodyDataInsertedFaces) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFacesResponseBodyDataInsertedFaces) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) SetFaceId(v string) *BatchAddFacesResponseBodyDataInsertedFaces {
	s.FaceId = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) SetImageURL(v string) *BatchAddFacesResponseBodyDataInsertedFaces {
	s.ImageURL = &v
	return s
}

func (s *BatchAddFacesResponseBodyDataInsertedFaces) SetQualitieScore(v float32) *BatchAddFacesResponseBodyDataInsertedFaces {
	s.QualitieScore = &v
	return s
}

type BatchAddFacesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchAddFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchAddFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFacesResponse) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponse) SetHeaders(v map[string]*string) *BatchAddFacesResponse {
	s.Headers = v
	return s
}

func (s *BatchAddFacesResponse) SetStatusCode(v int32) *BatchAddFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddFacesResponse) SetBody(v *BatchAddFacesResponseBody) *BatchAddFacesResponse {
	s.Body = v
	return s
}

type BeautifyBodyRequest struct {
	AgeRange            *BeautifyBodyRequestAgeRange    `json:"AgeRange,omitempty" xml:"AgeRange,omitempty" type:"Struct"`
	BodyBoxes           []*BeautifyBodyRequestBodyBoxes `json:"BodyBoxes,omitempty" xml:"BodyBoxes,omitempty" type:"Repeated"`
	Custom              *int64                          `json:"Custom,omitempty" xml:"Custom,omitempty"`
	FaceList            []*BeautifyBodyRequestFaceList  `json:"FaceList,omitempty" xml:"FaceList,omitempty" type:"Repeated"`
	FemaleLiquifyDegree *float32                        `json:"FemaleLiquifyDegree,omitempty" xml:"FemaleLiquifyDegree,omitempty"`
	ImageURL            *string                         `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	IsPregnant          *bool                           `json:"IsPregnant,omitempty" xml:"IsPregnant,omitempty"`
	LengthenDegree      *float32                        `json:"LengthenDegree,omitempty" xml:"LengthenDegree,omitempty"`
	MaleLiquifyDegree   *float32                        `json:"MaleLiquifyDegree,omitempty" xml:"MaleLiquifyDegree,omitempty"`
	OriginalHeight      *int64                          `json:"OriginalHeight,omitempty" xml:"OriginalHeight,omitempty"`
	OriginalWidth       *int64                          `json:"OriginalWidth,omitempty" xml:"OriginalWidth,omitempty"`
	PoseList            []*BeautifyBodyRequestPoseList  `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
}

func (s BeautifyBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyRequest) GoString() string {
	return s.String()
}

func (s *BeautifyBodyRequest) SetAgeRange(v *BeautifyBodyRequestAgeRange) *BeautifyBodyRequest {
	s.AgeRange = v
	return s
}

func (s *BeautifyBodyRequest) SetBodyBoxes(v []*BeautifyBodyRequestBodyBoxes) *BeautifyBodyRequest {
	s.BodyBoxes = v
	return s
}

func (s *BeautifyBodyRequest) SetCustom(v int64) *BeautifyBodyRequest {
	s.Custom = &v
	return s
}

func (s *BeautifyBodyRequest) SetFaceList(v []*BeautifyBodyRequestFaceList) *BeautifyBodyRequest {
	s.FaceList = v
	return s
}

func (s *BeautifyBodyRequest) SetFemaleLiquifyDegree(v float32) *BeautifyBodyRequest {
	s.FemaleLiquifyDegree = &v
	return s
}

func (s *BeautifyBodyRequest) SetImageURL(v string) *BeautifyBodyRequest {
	s.ImageURL = &v
	return s
}

func (s *BeautifyBodyRequest) SetIsPregnant(v bool) *BeautifyBodyRequest {
	s.IsPregnant = &v
	return s
}

func (s *BeautifyBodyRequest) SetLengthenDegree(v float32) *BeautifyBodyRequest {
	s.LengthenDegree = &v
	return s
}

func (s *BeautifyBodyRequest) SetMaleLiquifyDegree(v float32) *BeautifyBodyRequest {
	s.MaleLiquifyDegree = &v
	return s
}

func (s *BeautifyBodyRequest) SetOriginalHeight(v int64) *BeautifyBodyRequest {
	s.OriginalHeight = &v
	return s
}

func (s *BeautifyBodyRequest) SetOriginalWidth(v int64) *BeautifyBodyRequest {
	s.OriginalWidth = &v
	return s
}

func (s *BeautifyBodyRequest) SetPoseList(v []*BeautifyBodyRequestPoseList) *BeautifyBodyRequest {
	s.PoseList = v
	return s
}

type BeautifyBodyRequestAgeRange struct {
	AgeMax     *int64 `json:"AgeMax,omitempty" xml:"AgeMax,omitempty"`
	AgeMinimum *int64 `json:"AgeMinimum,omitempty" xml:"AgeMinimum,omitempty"`
}

func (s BeautifyBodyRequestAgeRange) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyRequestAgeRange) GoString() string {
	return s.String()
}

func (s *BeautifyBodyRequestAgeRange) SetAgeMax(v int64) *BeautifyBodyRequestAgeRange {
	s.AgeMax = &v
	return s
}

func (s *BeautifyBodyRequestAgeRange) SetAgeMinimum(v int64) *BeautifyBodyRequestAgeRange {
	s.AgeMinimum = &v
	return s
}

type BeautifyBodyRequestBodyBoxes struct {
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s BeautifyBodyRequestBodyBoxes) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyRequestBodyBoxes) GoString() string {
	return s.String()
}

func (s *BeautifyBodyRequestBodyBoxes) SetHeight(v float32) *BeautifyBodyRequestBodyBoxes {
	s.Height = &v
	return s
}

func (s *BeautifyBodyRequestBodyBoxes) SetWidth(v float32) *BeautifyBodyRequestBodyBoxes {
	s.Width = &v
	return s
}

func (s *BeautifyBodyRequestBodyBoxes) SetX(v float32) *BeautifyBodyRequestBodyBoxes {
	s.X = &v
	return s
}

func (s *BeautifyBodyRequestBodyBoxes) SetY(v float32) *BeautifyBodyRequestBodyBoxes {
	s.Y = &v
	return s
}

type BeautifyBodyRequestFaceList struct {
	Age     *int64                              `json:"Age,omitempty" xml:"Age,omitempty"`
	FaceBox *BeautifyBodyRequestFaceListFaceBox `json:"FaceBox,omitempty" xml:"FaceBox,omitempty" type:"Struct"`
	Gender  *int64                              `json:"Gender,omitempty" xml:"Gender,omitempty"`
}

func (s BeautifyBodyRequestFaceList) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyRequestFaceList) GoString() string {
	return s.String()
}

func (s *BeautifyBodyRequestFaceList) SetAge(v int64) *BeautifyBodyRequestFaceList {
	s.Age = &v
	return s
}

func (s *BeautifyBodyRequestFaceList) SetFaceBox(v *BeautifyBodyRequestFaceListFaceBox) *BeautifyBodyRequestFaceList {
	s.FaceBox = v
	return s
}

func (s *BeautifyBodyRequestFaceList) SetGender(v int64) *BeautifyBodyRequestFaceList {
	s.Gender = &v
	return s
}

type BeautifyBodyRequestFaceListFaceBox struct {
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s BeautifyBodyRequestFaceListFaceBox) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyRequestFaceListFaceBox) GoString() string {
	return s.String()
}

func (s *BeautifyBodyRequestFaceListFaceBox) SetHeight(v float32) *BeautifyBodyRequestFaceListFaceBox {
	s.Height = &v
	return s
}

func (s *BeautifyBodyRequestFaceListFaceBox) SetWidth(v float32) *BeautifyBodyRequestFaceListFaceBox {
	s.Width = &v
	return s
}

func (s *BeautifyBodyRequestFaceListFaceBox) SetX(v float32) *BeautifyBodyRequestFaceListFaceBox {
	s.X = &v
	return s
}

func (s *BeautifyBodyRequestFaceListFaceBox) SetY(v float32) *BeautifyBodyRequestFaceListFaceBox {
	s.Y = &v
	return s
}

type BeautifyBodyRequestPoseList struct {
	Pose []*BeautifyBodyRequestPoseListPose `json:"Pose,omitempty" xml:"Pose,omitempty" type:"Repeated"`
}

func (s BeautifyBodyRequestPoseList) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyRequestPoseList) GoString() string {
	return s.String()
}

func (s *BeautifyBodyRequestPoseList) SetPose(v []*BeautifyBodyRequestPoseListPose) *BeautifyBodyRequestPoseList {
	s.Pose = v
	return s
}

type BeautifyBodyRequestPoseListPose struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	X     *int64   `json:"X,omitempty" xml:"X,omitempty"`
	Y     *int64   `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s BeautifyBodyRequestPoseListPose) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyRequestPoseListPose) GoString() string {
	return s.String()
}

func (s *BeautifyBodyRequestPoseListPose) SetScore(v float32) *BeautifyBodyRequestPoseListPose {
	s.Score = &v
	return s
}

func (s *BeautifyBodyRequestPoseListPose) SetX(v int64) *BeautifyBodyRequestPoseListPose {
	s.X = &v
	return s
}

func (s *BeautifyBodyRequestPoseListPose) SetY(v int64) *BeautifyBodyRequestPoseListPose {
	s.Y = &v
	return s
}

type BeautifyBodyAdvanceRequest struct {
	AgeRange            *BeautifyBodyAdvanceRequestAgeRange    `json:"AgeRange,omitempty" xml:"AgeRange,omitempty" type:"Struct"`
	BodyBoxes           []*BeautifyBodyAdvanceRequestBodyBoxes `json:"BodyBoxes,omitempty" xml:"BodyBoxes,omitempty" type:"Repeated"`
	Custom              *int64                                 `json:"Custom,omitempty" xml:"Custom,omitempty"`
	FaceList            []*BeautifyBodyAdvanceRequestFaceList  `json:"FaceList,omitempty" xml:"FaceList,omitempty" type:"Repeated"`
	FemaleLiquifyDegree *float32                               `json:"FemaleLiquifyDegree,omitempty" xml:"FemaleLiquifyDegree,omitempty"`
	ImageURLObject      io.Reader                              `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	IsPregnant          *bool                                  `json:"IsPregnant,omitempty" xml:"IsPregnant,omitempty"`
	LengthenDegree      *float32                               `json:"LengthenDegree,omitempty" xml:"LengthenDegree,omitempty"`
	MaleLiquifyDegree   *float32                               `json:"MaleLiquifyDegree,omitempty" xml:"MaleLiquifyDegree,omitempty"`
	OriginalHeight      *int64                                 `json:"OriginalHeight,omitempty" xml:"OriginalHeight,omitempty"`
	OriginalWidth       *int64                                 `json:"OriginalWidth,omitempty" xml:"OriginalWidth,omitempty"`
	PoseList            []*BeautifyBodyAdvanceRequestPoseList  `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
}

func (s BeautifyBodyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BeautifyBodyAdvanceRequest) SetAgeRange(v *BeautifyBodyAdvanceRequestAgeRange) *BeautifyBodyAdvanceRequest {
	s.AgeRange = v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetBodyBoxes(v []*BeautifyBodyAdvanceRequestBodyBoxes) *BeautifyBodyAdvanceRequest {
	s.BodyBoxes = v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetCustom(v int64) *BeautifyBodyAdvanceRequest {
	s.Custom = &v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetFaceList(v []*BeautifyBodyAdvanceRequestFaceList) *BeautifyBodyAdvanceRequest {
	s.FaceList = v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetFemaleLiquifyDegree(v float32) *BeautifyBodyAdvanceRequest {
	s.FemaleLiquifyDegree = &v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetImageURLObject(v io.Reader) *BeautifyBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetIsPregnant(v bool) *BeautifyBodyAdvanceRequest {
	s.IsPregnant = &v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetLengthenDegree(v float32) *BeautifyBodyAdvanceRequest {
	s.LengthenDegree = &v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetMaleLiquifyDegree(v float32) *BeautifyBodyAdvanceRequest {
	s.MaleLiquifyDegree = &v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetOriginalHeight(v int64) *BeautifyBodyAdvanceRequest {
	s.OriginalHeight = &v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetOriginalWidth(v int64) *BeautifyBodyAdvanceRequest {
	s.OriginalWidth = &v
	return s
}

func (s *BeautifyBodyAdvanceRequest) SetPoseList(v []*BeautifyBodyAdvanceRequestPoseList) *BeautifyBodyAdvanceRequest {
	s.PoseList = v
	return s
}

type BeautifyBodyAdvanceRequestAgeRange struct {
	AgeMax     *int64 `json:"AgeMax,omitempty" xml:"AgeMax,omitempty"`
	AgeMinimum *int64 `json:"AgeMinimum,omitempty" xml:"AgeMinimum,omitempty"`
}

func (s BeautifyBodyAdvanceRequestAgeRange) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyAdvanceRequestAgeRange) GoString() string {
	return s.String()
}

func (s *BeautifyBodyAdvanceRequestAgeRange) SetAgeMax(v int64) *BeautifyBodyAdvanceRequestAgeRange {
	s.AgeMax = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestAgeRange) SetAgeMinimum(v int64) *BeautifyBodyAdvanceRequestAgeRange {
	s.AgeMinimum = &v
	return s
}

type BeautifyBodyAdvanceRequestBodyBoxes struct {
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s BeautifyBodyAdvanceRequestBodyBoxes) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyAdvanceRequestBodyBoxes) GoString() string {
	return s.String()
}

func (s *BeautifyBodyAdvanceRequestBodyBoxes) SetHeight(v float32) *BeautifyBodyAdvanceRequestBodyBoxes {
	s.Height = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestBodyBoxes) SetWidth(v float32) *BeautifyBodyAdvanceRequestBodyBoxes {
	s.Width = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestBodyBoxes) SetX(v float32) *BeautifyBodyAdvanceRequestBodyBoxes {
	s.X = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestBodyBoxes) SetY(v float32) *BeautifyBodyAdvanceRequestBodyBoxes {
	s.Y = &v
	return s
}

type BeautifyBodyAdvanceRequestFaceList struct {
	Age     *int64                                     `json:"Age,omitempty" xml:"Age,omitempty"`
	FaceBox *BeautifyBodyAdvanceRequestFaceListFaceBox `json:"FaceBox,omitempty" xml:"FaceBox,omitempty" type:"Struct"`
	Gender  *int64                                     `json:"Gender,omitempty" xml:"Gender,omitempty"`
}

func (s BeautifyBodyAdvanceRequestFaceList) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyAdvanceRequestFaceList) GoString() string {
	return s.String()
}

func (s *BeautifyBodyAdvanceRequestFaceList) SetAge(v int64) *BeautifyBodyAdvanceRequestFaceList {
	s.Age = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestFaceList) SetFaceBox(v *BeautifyBodyAdvanceRequestFaceListFaceBox) *BeautifyBodyAdvanceRequestFaceList {
	s.FaceBox = v
	return s
}

func (s *BeautifyBodyAdvanceRequestFaceList) SetGender(v int64) *BeautifyBodyAdvanceRequestFaceList {
	s.Gender = &v
	return s
}

type BeautifyBodyAdvanceRequestFaceListFaceBox struct {
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s BeautifyBodyAdvanceRequestFaceListFaceBox) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyAdvanceRequestFaceListFaceBox) GoString() string {
	return s.String()
}

func (s *BeautifyBodyAdvanceRequestFaceListFaceBox) SetHeight(v float32) *BeautifyBodyAdvanceRequestFaceListFaceBox {
	s.Height = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestFaceListFaceBox) SetWidth(v float32) *BeautifyBodyAdvanceRequestFaceListFaceBox {
	s.Width = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestFaceListFaceBox) SetX(v float32) *BeautifyBodyAdvanceRequestFaceListFaceBox {
	s.X = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestFaceListFaceBox) SetY(v float32) *BeautifyBodyAdvanceRequestFaceListFaceBox {
	s.Y = &v
	return s
}

type BeautifyBodyAdvanceRequestPoseList struct {
	Pose []*BeautifyBodyAdvanceRequestPoseListPose `json:"Pose,omitempty" xml:"Pose,omitempty" type:"Repeated"`
}

func (s BeautifyBodyAdvanceRequestPoseList) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyAdvanceRequestPoseList) GoString() string {
	return s.String()
}

func (s *BeautifyBodyAdvanceRequestPoseList) SetPose(v []*BeautifyBodyAdvanceRequestPoseListPose) *BeautifyBodyAdvanceRequestPoseList {
	s.Pose = v
	return s
}

type BeautifyBodyAdvanceRequestPoseListPose struct {
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	X     *int64   `json:"X,omitempty" xml:"X,omitempty"`
	Y     *int64   `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s BeautifyBodyAdvanceRequestPoseListPose) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyAdvanceRequestPoseListPose) GoString() string {
	return s.String()
}

func (s *BeautifyBodyAdvanceRequestPoseListPose) SetScore(v float32) *BeautifyBodyAdvanceRequestPoseListPose {
	s.Score = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestPoseListPose) SetX(v int64) *BeautifyBodyAdvanceRequestPoseListPose {
	s.X = &v
	return s
}

func (s *BeautifyBodyAdvanceRequestPoseListPose) SetY(v int64) *BeautifyBodyAdvanceRequestPoseListPose {
	s.Y = &v
	return s
}

type BeautifyBodyShrinkRequest struct {
	AgeRangeShrink      *string  `json:"AgeRange,omitempty" xml:"AgeRange,omitempty"`
	BodyBoxesShrink     *string  `json:"BodyBoxes,omitempty" xml:"BodyBoxes,omitempty"`
	Custom              *int64   `json:"Custom,omitempty" xml:"Custom,omitempty"`
	FaceListShrink      *string  `json:"FaceList,omitempty" xml:"FaceList,omitempty"`
	FemaleLiquifyDegree *float32 `json:"FemaleLiquifyDegree,omitempty" xml:"FemaleLiquifyDegree,omitempty"`
	ImageURL            *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	IsPregnant          *bool    `json:"IsPregnant,omitempty" xml:"IsPregnant,omitempty"`
	LengthenDegree      *float32 `json:"LengthenDegree,omitempty" xml:"LengthenDegree,omitempty"`
	MaleLiquifyDegree   *float32 `json:"MaleLiquifyDegree,omitempty" xml:"MaleLiquifyDegree,omitempty"`
	OriginalHeight      *int64   `json:"OriginalHeight,omitempty" xml:"OriginalHeight,omitempty"`
	OriginalWidth       *int64   `json:"OriginalWidth,omitempty" xml:"OriginalWidth,omitempty"`
	PoseListShrink      *string  `json:"PoseList,omitempty" xml:"PoseList,omitempty"`
}

func (s BeautifyBodyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyShrinkRequest) GoString() string {
	return s.String()
}

func (s *BeautifyBodyShrinkRequest) SetAgeRangeShrink(v string) *BeautifyBodyShrinkRequest {
	s.AgeRangeShrink = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetBodyBoxesShrink(v string) *BeautifyBodyShrinkRequest {
	s.BodyBoxesShrink = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetCustom(v int64) *BeautifyBodyShrinkRequest {
	s.Custom = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetFaceListShrink(v string) *BeautifyBodyShrinkRequest {
	s.FaceListShrink = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetFemaleLiquifyDegree(v float32) *BeautifyBodyShrinkRequest {
	s.FemaleLiquifyDegree = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetImageURL(v string) *BeautifyBodyShrinkRequest {
	s.ImageURL = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetIsPregnant(v bool) *BeautifyBodyShrinkRequest {
	s.IsPregnant = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetLengthenDegree(v float32) *BeautifyBodyShrinkRequest {
	s.LengthenDegree = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetMaleLiquifyDegree(v float32) *BeautifyBodyShrinkRequest {
	s.MaleLiquifyDegree = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetOriginalHeight(v int64) *BeautifyBodyShrinkRequest {
	s.OriginalHeight = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetOriginalWidth(v int64) *BeautifyBodyShrinkRequest {
	s.OriginalWidth = &v
	return s
}

func (s *BeautifyBodyShrinkRequest) SetPoseListShrink(v string) *BeautifyBodyShrinkRequest {
	s.PoseListShrink = &v
	return s
}

type BeautifyBodyResponseBody struct {
	// Id of the request
	Data      *BeautifyBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BeautifyBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyResponseBody) GoString() string {
	return s.String()
}

func (s *BeautifyBodyResponseBody) SetData(v *BeautifyBodyResponseBodyData) *BeautifyBodyResponseBody {
	s.Data = v
	return s
}

func (s *BeautifyBodyResponseBody) SetRequestId(v string) *BeautifyBodyResponseBody {
	s.RequestId = &v
	return s
}

type BeautifyBodyResponseBodyData struct {
	Action   *string `json:"Action,omitempty" xml:"Action,omitempty"`
	XFlowURL *string `json:"XFlowURL,omitempty" xml:"XFlowURL,omitempty"`
	YFlowURL *string `json:"YFlowURL,omitempty" xml:"YFlowURL,omitempty"`
}

func (s BeautifyBodyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *BeautifyBodyResponseBodyData) SetAction(v string) *BeautifyBodyResponseBodyData {
	s.Action = &v
	return s
}

func (s *BeautifyBodyResponseBodyData) SetXFlowURL(v string) *BeautifyBodyResponseBodyData {
	s.XFlowURL = &v
	return s
}

func (s *BeautifyBodyResponseBodyData) SetYFlowURL(v string) *BeautifyBodyResponseBodyData {
	s.YFlowURL = &v
	return s
}

type BeautifyBodyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BeautifyBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BeautifyBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s BeautifyBodyResponse) GoString() string {
	return s.String()
}

func (s *BeautifyBodyResponse) SetHeaders(v map[string]*string) *BeautifyBodyResponse {
	s.Headers = v
	return s
}

func (s *BeautifyBodyResponse) SetStatusCode(v int32) *BeautifyBodyResponse {
	s.StatusCode = &v
	return s
}

func (s *BeautifyBodyResponse) SetBody(v *BeautifyBodyResponseBody) *BeautifyBodyResponse {
	s.Body = v
	return s
}

type BlurFaceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BlurFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceRequest) GoString() string {
	return s.String()
}

func (s *BlurFaceRequest) SetImageURL(v string) *BlurFaceRequest {
	s.ImageURL = &v
	return s
}

type BlurFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BlurFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BlurFaceAdvanceRequest) SetImageURLObject(v io.Reader) *BlurFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type BlurFaceResponseBody struct {
	Data      *BlurFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BlurFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceResponseBody) GoString() string {
	return s.String()
}

func (s *BlurFaceResponseBody) SetData(v *BlurFaceResponseBodyData) *BlurFaceResponseBody {
	s.Data = v
	return s
}

func (s *BlurFaceResponseBody) SetRequestId(v string) *BlurFaceResponseBody {
	s.RequestId = &v
	return s
}

type BlurFaceResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BlurFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *BlurFaceResponseBodyData) SetImageURL(v string) *BlurFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

type BlurFaceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BlurFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BlurFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s BlurFaceResponse) GoString() string {
	return s.String()
}

func (s *BlurFaceResponse) SetHeaders(v map[string]*string) *BlurFaceResponse {
	s.Headers = v
	return s
}

func (s *BlurFaceResponse) SetStatusCode(v int32) *BlurFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *BlurFaceResponse) SetBody(v *BlurFaceResponseBody) *BlurFaceResponse {
	s.Body = v
	return s
}

type BodyPostureRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BodyPostureRequest) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureRequest) GoString() string {
	return s.String()
}

func (s *BodyPostureRequest) SetImageURL(v string) *BodyPostureRequest {
	s.ImageURL = &v
	return s
}

type BodyPostureAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BodyPostureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BodyPostureAdvanceRequest) SetImageURLObject(v io.Reader) *BodyPostureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type BodyPostureResponseBody struct {
	Data      *BodyPostureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BodyPostureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBody) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBody) SetData(v *BodyPostureResponseBodyData) *BodyPostureResponseBody {
	s.Data = v
	return s
}

func (s *BodyPostureResponseBody) SetRequestId(v string) *BodyPostureResponseBody {
	s.RequestId = &v
	return s
}

type BodyPostureResponseBodyData struct {
	MetaObject *BodyPostureResponseBodyDataMetaObject `json:"MetaObject,omitempty" xml:"MetaObject,omitempty" type:"Struct"`
	Outputs    []*BodyPostureResponseBodyDataOutputs  `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyData) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyData) SetMetaObject(v *BodyPostureResponseBodyDataMetaObject) *BodyPostureResponseBodyData {
	s.MetaObject = v
	return s
}

func (s *BodyPostureResponseBodyData) SetOutputs(v []*BodyPostureResponseBodyDataOutputs) *BodyPostureResponseBodyData {
	s.Outputs = v
	return s
}

type BodyPostureResponseBodyDataMetaObject struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BodyPostureResponseBodyDataMetaObject) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataMetaObject) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataMetaObject) SetHeight(v int32) *BodyPostureResponseBodyDataMetaObject {
	s.Height = &v
	return s
}

func (s *BodyPostureResponseBodyDataMetaObject) SetWidth(v int32) *BodyPostureResponseBodyDataMetaObject {
	s.Width = &v
	return s
}

type BodyPostureResponseBodyDataOutputs struct {
	HumanCount *int32                                       `json:"HumanCount,omitempty" xml:"HumanCount,omitempty"`
	Results    []*BodyPostureResponseBodyDataOutputsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputs) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputs) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputs) SetHumanCount(v int32) *BodyPostureResponseBodyDataOutputs {
	s.HumanCount = &v
	return s
}

func (s *BodyPostureResponseBodyDataOutputs) SetResults(v []*BodyPostureResponseBodyDataOutputsResults) *BodyPostureResponseBodyDataOutputs {
	s.Results = v
	return s
}

type BodyPostureResponseBodyDataOutputsResults struct {
	Bodies []*BodyPostureResponseBodyDataOutputsResultsBodies `json:"Bodies,omitempty" xml:"Bodies,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputsResults) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputsResults) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputsResults) SetBodies(v []*BodyPostureResponseBodyDataOutputsResultsBodies) *BodyPostureResponseBodyDataOutputsResults {
	s.Bodies = v
	return s
}

type BodyPostureResponseBodyDataOutputsResultsBodies struct {
	Confident *float32                                                    `json:"Confident,omitempty" xml:"Confident,omitempty"`
	Label     *string                                                     `json:"Label,omitempty" xml:"Label,omitempty"`
	Positions []*BodyPostureResponseBodyDataOutputsResultsBodiesPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputsResultsBodies) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputsResultsBodies) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) SetConfident(v float32) *BodyPostureResponseBodyDataOutputsResultsBodies {
	s.Confident = &v
	return s
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) SetLabel(v string) *BodyPostureResponseBodyDataOutputsResultsBodies {
	s.Label = &v
	return s
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) SetPositions(v []*BodyPostureResponseBodyDataOutputsResultsBodiesPositions) *BodyPostureResponseBodyDataOutputsResultsBodies {
	s.Positions = v
	return s
}

type BodyPostureResponseBodyDataOutputsResultsBodiesPositions struct {
	Points []*float32 `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputsResultsBodiesPositions) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputsResultsBodiesPositions) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodiesPositions) SetPoints(v []*float32) *BodyPostureResponseBodyDataOutputsResultsBodiesPositions {
	s.Points = v
	return s
}

type BodyPostureResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BodyPostureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BodyPostureResponse) String() string {
	return tea.Prettify(s)
}

func (s BodyPostureResponse) GoString() string {
	return s.String()
}

func (s *BodyPostureResponse) SetHeaders(v map[string]*string) *BodyPostureResponse {
	s.Headers = v
	return s
}

func (s *BodyPostureResponse) SetStatusCode(v int32) *BodyPostureResponse {
	s.StatusCode = &v
	return s
}

func (s *BodyPostureResponse) SetBody(v *BodyPostureResponseBody) *BodyPostureResponse {
	s.Body = v
	return s
}

type CompareFaceRequest struct {
	ImageDataA            *string  `json:"ImageDataA,omitempty" xml:"ImageDataA,omitempty"`
	ImageDataB            *string  `json:"ImageDataB,omitempty" xml:"ImageDataB,omitempty"`
	ImageURLA             *string  `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	ImageURLB             *string  `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
	QualityScoreThreshold *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
}

func (s CompareFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceRequest) SetImageDataA(v string) *CompareFaceRequest {
	s.ImageDataA = &v
	return s
}

func (s *CompareFaceRequest) SetImageDataB(v string) *CompareFaceRequest {
	s.ImageDataB = &v
	return s
}

func (s *CompareFaceRequest) SetImageURLA(v string) *CompareFaceRequest {
	s.ImageURLA = &v
	return s
}

func (s *CompareFaceRequest) SetImageURLB(v string) *CompareFaceRequest {
	s.ImageURLB = &v
	return s
}

func (s *CompareFaceRequest) SetQualityScoreThreshold(v float32) *CompareFaceRequest {
	s.QualityScoreThreshold = &v
	return s
}

type CompareFaceAdvanceRequest struct {
	ImageDataA            *string   `json:"ImageDataA,omitempty" xml:"ImageDataA,omitempty"`
	ImageDataB            *string   `json:"ImageDataB,omitempty" xml:"ImageDataB,omitempty"`
	ImageURLAObject       io.Reader `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	ImageURLBObject       io.Reader `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
	QualityScoreThreshold *float32  `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
}

func (s CompareFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceAdvanceRequest) SetImageDataA(v string) *CompareFaceAdvanceRequest {
	s.ImageDataA = &v
	return s
}

func (s *CompareFaceAdvanceRequest) SetImageDataB(v string) *CompareFaceAdvanceRequest {
	s.ImageDataB = &v
	return s
}

func (s *CompareFaceAdvanceRequest) SetImageURLAObject(v io.Reader) *CompareFaceAdvanceRequest {
	s.ImageURLAObject = v
	return s
}

func (s *CompareFaceAdvanceRequest) SetImageURLBObject(v io.Reader) *CompareFaceAdvanceRequest {
	s.ImageURLBObject = v
	return s
}

func (s *CompareFaceAdvanceRequest) SetQualityScoreThreshold(v float32) *CompareFaceAdvanceRequest {
	s.QualityScoreThreshold = &v
	return s
}

type CompareFaceResponseBody struct {
	Data      *CompareFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompareFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFaceResponseBody) SetData(v *CompareFaceResponseBodyData) *CompareFaceResponseBody {
	s.Data = v
	return s
}

func (s *CompareFaceResponseBody) SetRequestId(v string) *CompareFaceResponseBody {
	s.RequestId = &v
	return s
}

type CompareFaceResponseBodyData struct {
	Confidence    *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	MessageTips   *string  `json:"MessageTips,omitempty" xml:"MessageTips,omitempty"`
	QualityScoreA *float32 `json:"QualityScoreA,omitempty" xml:"QualityScoreA,omitempty"`
	QualityScoreB *float32 `json:"QualityScoreB,omitempty" xml:"QualityScoreB,omitempty"`
	// 1
	RectAList []*int32 `json:"RectAList,omitempty" xml:"RectAList,omitempty" type:"Repeated"`
	// 1
	RectBList []*int32 `json:"RectBList,omitempty" xml:"RectBList,omitempty" type:"Repeated"`
	// 1
	Thresholds []*float32 `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Repeated"`
}

func (s CompareFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompareFaceResponseBodyData) SetConfidence(v float32) *CompareFaceResponseBodyData {
	s.Confidence = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetMessageTips(v string) *CompareFaceResponseBodyData {
	s.MessageTips = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetQualityScoreA(v float32) *CompareFaceResponseBodyData {
	s.QualityScoreA = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetQualityScoreB(v float32) *CompareFaceResponseBodyData {
	s.QualityScoreB = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetRectAList(v []*int32) *CompareFaceResponseBodyData {
	s.RectAList = v
	return s
}

func (s *CompareFaceResponseBodyData) SetRectBList(v []*int32) *CompareFaceResponseBodyData {
	s.RectBList = v
	return s
}

func (s *CompareFaceResponseBodyData) SetThresholds(v []*float32) *CompareFaceResponseBodyData {
	s.Thresholds = v
	return s
}

type CompareFaceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompareFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompareFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareFaceResponse) GoString() string {
	return s.String()
}

func (s *CompareFaceResponse) SetHeaders(v map[string]*string) *CompareFaceResponse {
	s.Headers = v
	return s
}

func (s *CompareFaceResponse) SetStatusCode(v int32) *CompareFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareFaceResponse) SetBody(v *CompareFaceResponseBody) *CompareFaceResponse {
	s.Body = v
	return s
}

type CountCrowdRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	IsShow   *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
}

func (s CountCrowdRequest) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdRequest) GoString() string {
	return s.String()
}

func (s *CountCrowdRequest) SetImageURL(v string) *CountCrowdRequest {
	s.ImageURL = &v
	return s
}

func (s *CountCrowdRequest) SetIsShow(v bool) *CountCrowdRequest {
	s.IsShow = &v
	return s
}

type CountCrowdAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	IsShow         *bool     `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
}

func (s CountCrowdAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CountCrowdAdvanceRequest) SetImageURLObject(v io.Reader) *CountCrowdAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *CountCrowdAdvanceRequest) SetIsShow(v bool) *CountCrowdAdvanceRequest {
	s.IsShow = &v
	return s
}

type CountCrowdResponseBody struct {
	Data      *CountCrowdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CountCrowdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *CountCrowdResponseBody) SetData(v *CountCrowdResponseBodyData) *CountCrowdResponseBody {
	s.Data = v
	return s
}

func (s *CountCrowdResponseBody) SetRequestId(v string) *CountCrowdResponseBody {
	s.RequestId = &v
	return s
}

type CountCrowdResponseBodyData struct {
	HotMap       *string `json:"HotMap,omitempty" xml:"HotMap,omitempty"`
	PeopleNumber *int32  `json:"PeopleNumber,omitempty" xml:"PeopleNumber,omitempty"`
}

func (s CountCrowdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdResponseBodyData) GoString() string {
	return s.String()
}

func (s *CountCrowdResponseBodyData) SetHotMap(v string) *CountCrowdResponseBodyData {
	s.HotMap = &v
	return s
}

func (s *CountCrowdResponseBodyData) SetPeopleNumber(v int32) *CountCrowdResponseBodyData {
	s.PeopleNumber = &v
	return s
}

type CountCrowdResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountCrowdResponse) String() string {
	return tea.Prettify(s)
}

func (s CountCrowdResponse) GoString() string {
	return s.String()
}

func (s *CountCrowdResponse) SetHeaders(v map[string]*string) *CountCrowdResponse {
	s.Headers = v
	return s
}

func (s *CountCrowdResponse) SetStatusCode(v int32) *CountCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *CountCrowdResponse) SetBody(v *CountCrowdResponseBody) *CountCrowdResponse {
	s.Body = v
	return s
}

type CreateFaceDbRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateFaceDbRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceDbRequest) GoString() string {
	return s.String()
}

func (s *CreateFaceDbRequest) SetName(v string) *CreateFaceDbRequest {
	s.Name = &v
	return s
}

type CreateFaceDbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFaceDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceDbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFaceDbResponseBody) SetRequestId(v string) *CreateFaceDbResponseBody {
	s.RequestId = &v
	return s
}

type CreateFaceDbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFaceDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFaceDbResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFaceDbResponse) GoString() string {
	return s.String()
}

func (s *CreateFaceDbResponse) SetHeaders(v map[string]*string) *CreateFaceDbResponse {
	s.Headers = v
	return s
}

func (s *CreateFaceDbResponse) SetStatusCode(v int32) *CreateFaceDbResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFaceDbResponse) SetBody(v *CreateFaceDbResponseBody) *CreateFaceDbResponse {
	s.Body = v
	return s
}

type DeleteFaceRequest struct {
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
}

func (s DeleteFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceRequest) SetDbName(v string) *DeleteFaceRequest {
	s.DbName = &v
	return s
}

func (s *DeleteFaceRequest) SetFaceId(v string) *DeleteFaceRequest {
	s.FaceId = &v
	return s
}

type DeleteFaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceResponseBody) SetRequestId(v string) *DeleteFaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceResponse) SetHeaders(v map[string]*string) *DeleteFaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceResponse) SetStatusCode(v int32) *DeleteFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceResponse) SetBody(v *DeleteFaceResponseBody) *DeleteFaceResponse {
	s.Body = v
	return s
}

type DeleteFaceDbRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteFaceDbRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDbRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceDbRequest) SetName(v string) *DeleteFaceDbRequest {
	s.Name = &v
	return s
}

type DeleteFaceDbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceDbResponseBody) SetRequestId(v string) *DeleteFaceDbResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceDbResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFaceDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceDbResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceDbResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceDbResponse) SetHeaders(v map[string]*string) *DeleteFaceDbResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceDbResponse) SetStatusCode(v int32) *DeleteFaceDbResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceDbResponse) SetBody(v *DeleteFaceDbResponseBody) *DeleteFaceDbResponse {
	s.Body = v
	return s
}

type DeleteFaceEntityRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s DeleteFaceEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceEntityRequest) SetDbName(v string) *DeleteFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *DeleteFaceEntityRequest) SetEntityId(v string) *DeleteFaceEntityRequest {
	s.EntityId = &v
	return s
}

type DeleteFaceEntityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceEntityResponseBody) SetRequestId(v string) *DeleteFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceEntityResponse) SetHeaders(v map[string]*string) *DeleteFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceEntityResponse) SetStatusCode(v int32) *DeleteFaceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceEntityResponse) SetBody(v *DeleteFaceEntityResponseBody) *DeleteFaceEntityResponse {
	s.Body = v
	return s
}

type DeleteFaceImageTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteFaceImageTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceImageTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageTemplateRequest) SetTemplateId(v string) *DeleteFaceImageTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteFaceImageTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceImageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceImageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageTemplateResponseBody) SetRequestId(v string) *DeleteFaceImageTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFaceImageTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFaceImageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFaceImageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFaceImageTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceImageTemplateResponse) SetHeaders(v map[string]*string) *DeleteFaceImageTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceImageTemplateResponse) SetStatusCode(v int32) *DeleteFaceImageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceImageTemplateResponse) SetBody(v *DeleteFaceImageTemplateResponseBody) *DeleteFaceImageTemplateResponse {
	s.Body = v
	return s
}

type DetectBodyCountRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectBodyCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountRequest) GoString() string {
	return s.String()
}

func (s *DetectBodyCountRequest) SetImageURL(v string) *DetectBodyCountRequest {
	s.ImageURL = &v
	return s
}

type DetectBodyCountAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectBodyCountAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectBodyCountAdvanceRequest) SetImageURLObject(v io.Reader) *DetectBodyCountAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectBodyCountResponseBody struct {
	Data      *DetectBodyCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectBodyCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DetectBodyCountResponseBody) SetData(v *DetectBodyCountResponseBodyData) *DetectBodyCountResponseBody {
	s.Data = v
	return s
}

func (s *DetectBodyCountResponseBody) SetRequestId(v string) *DetectBodyCountResponseBody {
	s.RequestId = &v
	return s
}

type DetectBodyCountResponseBodyData struct {
	PersonNumber *int32 `json:"PersonNumber,omitempty" xml:"PersonNumber,omitempty"`
}

func (s DetectBodyCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectBodyCountResponseBodyData) SetPersonNumber(v int32) *DetectBodyCountResponseBodyData {
	s.PersonNumber = &v
	return s
}

type DetectBodyCountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectBodyCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectBodyCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectBodyCountResponse) GoString() string {
	return s.String()
}

func (s *DetectBodyCountResponse) SetHeaders(v map[string]*string) *DetectBodyCountResponse {
	s.Headers = v
	return s
}

func (s *DetectBodyCountResponse) SetStatusCode(v int32) *DetectBodyCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectBodyCountResponse) SetBody(v *DetectBodyCountResponseBody) *DetectBodyCountResponse {
	s.Body = v
	return s
}

type DetectCelebrityRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectCelebrityRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityRequest) GoString() string {
	return s.String()
}

func (s *DetectCelebrityRequest) SetImageURL(v string) *DetectCelebrityRequest {
	s.ImageURL = &v
	return s
}

type DetectCelebrityAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectCelebrityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectCelebrityAdvanceRequest) SetImageURLObject(v io.Reader) *DetectCelebrityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectCelebrityResponseBody struct {
	Data      *DetectCelebrityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectCelebrityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityResponseBody) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponseBody) SetData(v *DetectCelebrityResponseBodyData) *DetectCelebrityResponseBody {
	s.Data = v
	return s
}

func (s *DetectCelebrityResponseBody) SetRequestId(v string) *DetectCelebrityResponseBody {
	s.RequestId = &v
	return s
}

type DetectCelebrityResponseBodyData struct {
	FaceRecognizeResults []*DetectCelebrityResponseBodyDataFaceRecognizeResults `json:"FaceRecognizeResults,omitempty" xml:"FaceRecognizeResults,omitempty" type:"Repeated"`
	Height               *int32                                                 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width                *int32                                                 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectCelebrityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponseBodyData) SetFaceRecognizeResults(v []*DetectCelebrityResponseBodyDataFaceRecognizeResults) *DetectCelebrityResponseBodyData {
	s.FaceRecognizeResults = v
	return s
}

func (s *DetectCelebrityResponseBodyData) SetHeight(v int32) *DetectCelebrityResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectCelebrityResponseBodyData) SetWidth(v int32) *DetectCelebrityResponseBodyData {
	s.Width = &v
	return s
}

type DetectCelebrityResponseBodyDataFaceRecognizeResults struct {
	// 1
	FaceBoxes []*float32 `json:"FaceBoxes,omitempty" xml:"FaceBoxes,omitempty" type:"Repeated"`
	Name      *string    `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DetectCelebrityResponseBodyDataFaceRecognizeResults) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityResponseBodyDataFaceRecognizeResults) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponseBodyDataFaceRecognizeResults) SetFaceBoxes(v []*float32) *DetectCelebrityResponseBodyDataFaceRecognizeResults {
	s.FaceBoxes = v
	return s
}

func (s *DetectCelebrityResponseBodyDataFaceRecognizeResults) SetName(v string) *DetectCelebrityResponseBodyDataFaceRecognizeResults {
	s.Name = &v
	return s
}

type DetectCelebrityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectCelebrityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectCelebrityResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectCelebrityResponse) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponse) SetHeaders(v map[string]*string) *DetectCelebrityResponse {
	s.Headers = v
	return s
}

func (s *DetectCelebrityResponse) SetStatusCode(v int32) *DetectCelebrityResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectCelebrityResponse) SetBody(v *DetectCelebrityResponseBody) *DetectCelebrityResponse {
	s.Body = v
	return s
}

type DetectChefCapRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectChefCapRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapRequest) GoString() string {
	return s.String()
}

func (s *DetectChefCapRequest) SetImageURL(v string) *DetectChefCapRequest {
	s.ImageURL = &v
	return s
}

type DetectChefCapAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectChefCapAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectChefCapAdvanceRequest) SetImageURLObject(v io.Reader) *DetectChefCapAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectChefCapResponseBody struct {
	Data      *DetectChefCapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectChefCapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapResponseBody) GoString() string {
	return s.String()
}

func (s *DetectChefCapResponseBody) SetData(v *DetectChefCapResponseBodyData) *DetectChefCapResponseBody {
	s.Data = v
	return s
}

func (s *DetectChefCapResponseBody) SetRequestId(v string) *DetectChefCapResponseBody {
	s.RequestId = &v
	return s
}

type DetectChefCapResponseBodyData struct {
	Elements []*DetectChefCapResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectChefCapResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectChefCapResponseBodyData) SetElements(v []*DetectChefCapResponseBodyDataElements) *DetectChefCapResponseBodyData {
	s.Elements = v
	return s
}

type DetectChefCapResponseBodyDataElements struct {
	Box        []*float32 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
	Category   *string    `json:"Category,omitempty" xml:"Category,omitempty"`
	Confidence *float32   `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s DetectChefCapResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectChefCapResponseBodyDataElements) SetBox(v []*float32) *DetectChefCapResponseBodyDataElements {
	s.Box = v
	return s
}

func (s *DetectChefCapResponseBodyDataElements) SetCategory(v string) *DetectChefCapResponseBodyDataElements {
	s.Category = &v
	return s
}

func (s *DetectChefCapResponseBodyDataElements) SetConfidence(v float32) *DetectChefCapResponseBodyDataElements {
	s.Confidence = &v
	return s
}

type DetectChefCapResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectChefCapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectChefCapResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectChefCapResponse) GoString() string {
	return s.String()
}

func (s *DetectChefCapResponse) SetHeaders(v map[string]*string) *DetectChefCapResponse {
	s.Headers = v
	return s
}

func (s *DetectChefCapResponse) SetStatusCode(v int32) *DetectChefCapResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectChefCapResponse) SetBody(v *DetectChefCapResponseBody) *DetectChefCapResponse {
	s.Body = v
	return s
}

type DetectFaceRequest struct {
	ImageURL      *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Landmark      *bool   `json:"Landmark,omitempty" xml:"Landmark,omitempty"`
	MaxFaceNumber *int64  `json:"MaxFaceNumber,omitempty" xml:"MaxFaceNumber,omitempty"`
	Pose          *bool   `json:"Pose,omitempty" xml:"Pose,omitempty"`
	Quality       *bool   `json:"Quality,omitempty" xml:"Quality,omitempty"`
}

func (s DetectFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceRequest) SetImageURL(v string) *DetectFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectFaceRequest) SetLandmark(v bool) *DetectFaceRequest {
	s.Landmark = &v
	return s
}

func (s *DetectFaceRequest) SetMaxFaceNumber(v int64) *DetectFaceRequest {
	s.MaxFaceNumber = &v
	return s
}

func (s *DetectFaceRequest) SetPose(v bool) *DetectFaceRequest {
	s.Pose = &v
	return s
}

func (s *DetectFaceRequest) SetQuality(v bool) *DetectFaceRequest {
	s.Quality = &v
	return s
}

type DetectFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Landmark       *bool     `json:"Landmark,omitempty" xml:"Landmark,omitempty"`
	MaxFaceNumber  *int64    `json:"MaxFaceNumber,omitempty" xml:"MaxFaceNumber,omitempty"`
	Pose           *bool     `json:"Pose,omitempty" xml:"Pose,omitempty"`
	Quality        *bool     `json:"Quality,omitempty" xml:"Quality,omitempty"`
}

func (s DetectFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAdvanceRequest) SetImageURLObject(v io.Reader) *DetectFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectFaceAdvanceRequest) SetLandmark(v bool) *DetectFaceAdvanceRequest {
	s.Landmark = &v
	return s
}

func (s *DetectFaceAdvanceRequest) SetMaxFaceNumber(v int64) *DetectFaceAdvanceRequest {
	s.MaxFaceNumber = &v
	return s
}

func (s *DetectFaceAdvanceRequest) SetPose(v bool) *DetectFaceAdvanceRequest {
	s.Pose = &v
	return s
}

func (s *DetectFaceAdvanceRequest) SetQuality(v bool) *DetectFaceAdvanceRequest {
	s.Quality = &v
	return s
}

type DetectFaceResponseBody struct {
	Data      *DetectFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBody) SetData(v *DetectFaceResponseBodyData) *DetectFaceResponseBody {
	s.Data = v
	return s
}

func (s *DetectFaceResponseBody) SetRequestId(v string) *DetectFaceResponseBody {
	s.RequestId = &v
	return s
}

type DetectFaceResponseBodyData struct {
	FaceCount *int32 `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	// 1
	FaceProbabilityList []*float32 `json:"FaceProbabilityList,omitempty" xml:"FaceProbabilityList,omitempty" type:"Repeated"`
	// 1
	FaceRectangles []*int32 `json:"FaceRectangles,omitempty" xml:"FaceRectangles,omitempty" type:"Repeated"`
	LandmarkCount  *int32   `json:"LandmarkCount,omitempty" xml:"LandmarkCount,omitempty"`
	// 1
	Landmarks []*float32 `json:"Landmarks,omitempty" xml:"Landmarks,omitempty" type:"Repeated"`
	// 1
	PoseList []*float32 `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
	// 1
	Pupils    []*float32                           `json:"Pupils,omitempty" xml:"Pupils,omitempty" type:"Repeated"`
	Qualities *DetectFaceResponseBodyDataQualities `json:"Qualities,omitempty" xml:"Qualities,omitempty" type:"Struct"`
}

func (s DetectFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBodyData) SetFaceCount(v int32) *DetectFaceResponseBodyData {
	s.FaceCount = &v
	return s
}

func (s *DetectFaceResponseBodyData) SetFaceProbabilityList(v []*float32) *DetectFaceResponseBodyData {
	s.FaceProbabilityList = v
	return s
}

func (s *DetectFaceResponseBodyData) SetFaceRectangles(v []*int32) *DetectFaceResponseBodyData {
	s.FaceRectangles = v
	return s
}

func (s *DetectFaceResponseBodyData) SetLandmarkCount(v int32) *DetectFaceResponseBodyData {
	s.LandmarkCount = &v
	return s
}

func (s *DetectFaceResponseBodyData) SetLandmarks(v []*float32) *DetectFaceResponseBodyData {
	s.Landmarks = v
	return s
}

func (s *DetectFaceResponseBodyData) SetPoseList(v []*float32) *DetectFaceResponseBodyData {
	s.PoseList = v
	return s
}

func (s *DetectFaceResponseBodyData) SetPupils(v []*float32) *DetectFaceResponseBodyData {
	s.Pupils = v
	return s
}

func (s *DetectFaceResponseBodyData) SetQualities(v *DetectFaceResponseBodyDataQualities) *DetectFaceResponseBodyData {
	s.Qualities = v
	return s
}

type DetectFaceResponseBodyDataQualities struct {
	// 1
	BlurList []*float32 `json:"BlurList,omitempty" xml:"BlurList,omitempty" type:"Repeated"`
	// 1
	FnfList []*float32 `json:"FnfList,omitempty" xml:"FnfList,omitempty" type:"Repeated"`
	// 1
	GlassList []*float32 `json:"GlassList,omitempty" xml:"GlassList,omitempty" type:"Repeated"`
	// 1
	IlluList []*float32 `json:"IlluList,omitempty" xml:"IlluList,omitempty" type:"Repeated"`
	// 1
	MaskList []*float32 `json:"MaskList,omitempty" xml:"MaskList,omitempty" type:"Repeated"`
	// 1
	NoiseList []*float32 `json:"NoiseList,omitempty" xml:"NoiseList,omitempty" type:"Repeated"`
	// 1
	PoseList []*float32 `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
	// 1
	ScoreList []*float32 `json:"ScoreList,omitempty" xml:"ScoreList,omitempty" type:"Repeated"`
}

func (s DetectFaceResponseBodyDataQualities) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponseBodyDataQualities) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBodyDataQualities) SetBlurList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.BlurList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetFnfList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.FnfList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetGlassList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.GlassList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetIlluList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.IlluList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetMaskList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.MaskList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetNoiseList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.NoiseList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetPoseList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.PoseList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetScoreList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.ScoreList = v
	return s
}

type DetectFaceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceResponse) SetHeaders(v map[string]*string) *DetectFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectFaceResponse) SetStatusCode(v int32) *DetectFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectFaceResponse) SetBody(v *DetectFaceResponseBody) *DetectFaceResponse {
	s.Body = v
	return s
}

type DetectIPCPedestrianRequest struct {
	Height    *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Width     *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectIPCPedestrianRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianRequest) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianRequest) SetHeight(v int32) *DetectIPCPedestrianRequest {
	s.Height = &v
	return s
}

func (s *DetectIPCPedestrianRequest) SetImageData(v string) *DetectIPCPedestrianRequest {
	s.ImageData = &v
	return s
}

func (s *DetectIPCPedestrianRequest) SetImageURL(v string) *DetectIPCPedestrianRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectIPCPedestrianRequest) SetWidth(v int32) *DetectIPCPedestrianRequest {
	s.Width = &v
	return s
}

type DetectIPCPedestrianAdvanceRequest struct {
	Height         *int32    `json:"Height,omitempty" xml:"Height,omitempty"`
	ImageData      *string   `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Width          *int32    `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectIPCPedestrianAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianAdvanceRequest) SetHeight(v int32) *DetectIPCPedestrianAdvanceRequest {
	s.Height = &v
	return s
}

func (s *DetectIPCPedestrianAdvanceRequest) SetImageData(v string) *DetectIPCPedestrianAdvanceRequest {
	s.ImageData = &v
	return s
}

func (s *DetectIPCPedestrianAdvanceRequest) SetImageURLObject(v io.Reader) *DetectIPCPedestrianAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectIPCPedestrianAdvanceRequest) SetWidth(v int32) *DetectIPCPedestrianAdvanceRequest {
	s.Width = &v
	return s
}

type DetectIPCPedestrianResponseBody struct {
	Data *DetectIPCPedestrianResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectIPCPedestrianResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponseBody) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponseBody) SetData(v *DetectIPCPedestrianResponseBodyData) *DetectIPCPedestrianResponseBody {
	s.Data = v
	return s
}

func (s *DetectIPCPedestrianResponseBody) SetRequestId(v string) *DetectIPCPedestrianResponseBody {
	s.RequestId = &v
	return s
}

type DetectIPCPedestrianResponseBodyData struct {
	ImageInfoList []*DetectIPCPedestrianResponseBodyDataImageInfoList `json:"ImageInfoList,omitempty" xml:"ImageInfoList,omitempty" type:"Repeated"`
}

func (s DetectIPCPedestrianResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponseBodyData) SetImageInfoList(v []*DetectIPCPedestrianResponseBodyDataImageInfoList) *DetectIPCPedestrianResponseBodyData {
	s.ImageInfoList = v
	return s
}

type DetectIPCPedestrianResponseBodyDataImageInfoList struct {
	Elements []*DetectIPCPedestrianResponseBodyDataImageInfoListElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectIPCPedestrianResponseBodyDataImageInfoList) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponseBodyDataImageInfoList) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponseBodyDataImageInfoList) SetElements(v []*DetectIPCPedestrianResponseBodyDataImageInfoListElements) *DetectIPCPedestrianResponseBodyDataImageInfoList {
	s.Elements = v
	return s
}

type DetectIPCPedestrianResponseBodyDataImageInfoListElements struct {
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s DetectIPCPedestrianResponseBodyDataImageInfoListElements) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponseBodyDataImageInfoListElements) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponseBodyDataImageInfoListElements) SetBoxes(v []*int32) *DetectIPCPedestrianResponseBodyDataImageInfoListElements {
	s.Boxes = v
	return s
}

func (s *DetectIPCPedestrianResponseBodyDataImageInfoListElements) SetScore(v float32) *DetectIPCPedestrianResponseBodyDataImageInfoListElements {
	s.Score = &v
	return s
}

type DetectIPCPedestrianResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectIPCPedestrianResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectIPCPedestrianResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianResponse) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianResponse) SetHeaders(v map[string]*string) *DetectIPCPedestrianResponse {
	s.Headers = v
	return s
}

func (s *DetectIPCPedestrianResponse) SetStatusCode(v int32) *DetectIPCPedestrianResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectIPCPedestrianResponse) SetBody(v *DetectIPCPedestrianResponseBody) *DetectIPCPedestrianResponse {
	s.Body = v
	return s
}

type DetectLivingFaceRequest struct {
	Tasks []*DetectLivingFaceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DetectLivingFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceRequest) SetTasks(v []*DetectLivingFaceRequestTasks) *DetectLivingFaceRequest {
	s.Tasks = v
	return s
}

type DetectLivingFaceRequestTasks struct {
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectLivingFaceRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceRequestTasks) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceRequestTasks) SetImageData(v string) *DetectLivingFaceRequestTasks {
	s.ImageData = &v
	return s
}

func (s *DetectLivingFaceRequestTasks) SetImageURL(v string) *DetectLivingFaceRequestTasks {
	s.ImageURL = &v
	return s
}

type DetectLivingFaceAdvanceRequest struct {
	Tasks []*DetectLivingFaceAdvanceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DetectLivingFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceAdvanceRequest) SetTasks(v []*DetectLivingFaceAdvanceRequestTasks) *DetectLivingFaceAdvanceRequest {
	s.Tasks = v
	return s
}

type DetectLivingFaceAdvanceRequestTasks struct {
	ImageData      *string   `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectLivingFaceAdvanceRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceAdvanceRequestTasks) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceAdvanceRequestTasks) SetImageData(v string) *DetectLivingFaceAdvanceRequestTasks {
	s.ImageData = &v
	return s
}

func (s *DetectLivingFaceAdvanceRequestTasks) SetImageURLObject(v io.Reader) *DetectLivingFaceAdvanceRequestTasks {
	s.ImageURLObject = v
	return s
}

type DetectLivingFaceResponseBody struct {
	Data      *DetectLivingFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectLivingFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBody) SetData(v *DetectLivingFaceResponseBodyData) *DetectLivingFaceResponseBody {
	s.Data = v
	return s
}

func (s *DetectLivingFaceResponseBody) SetRequestId(v string) *DetectLivingFaceResponseBody {
	s.RequestId = &v
	return s
}

type DetectLivingFaceResponseBodyData struct {
	Elements []*DetectLivingFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectLivingFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyData) SetElements(v []*DetectLivingFaceResponseBodyDataElements) *DetectLivingFaceResponseBodyData {
	s.Elements = v
	return s
}

type DetectLivingFaceResponseBodyDataElements struct {
	FaceNumber *int64                                             `json:"FaceNumber,omitempty" xml:"FaceNumber,omitempty"`
	ImageURL   *string                                            `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Results    []*DetectLivingFaceResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	TaskId     *string                                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElements) SetFaceNumber(v int64) *DetectLivingFaceResponseBodyDataElements {
	s.FaceNumber = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElements) SetImageURL(v string) *DetectLivingFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElements) SetResults(v []*DetectLivingFaceResponseBodyDataElementsResults) *DetectLivingFaceResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElements) SetTaskId(v string) *DetectLivingFaceResponseBodyDataElements {
	s.TaskId = &v
	return s
}

type DetectLivingFaceResponseBodyDataElementsResults struct {
	Frames      []*DetectLivingFaceResponseBodyDataElementsResultsFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	Label       *string                                                  `json:"Label,omitempty" xml:"Label,omitempty"`
	MessageTips *string                                                  `json:"MessageTips,omitempty" xml:"MessageTips,omitempty"`
	Rate        *float32                                                 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Rect        *DetectLivingFaceResponseBodyDataElementsResultsRect     `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
	Suggestion  *string                                                  `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElementsResults) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetFrames(v []*DetectLivingFaceResponseBodyDataElementsResultsFrames) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Frames = v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetLabel(v string) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetMessageTips(v string) *DetectLivingFaceResponseBodyDataElementsResults {
	s.MessageTips = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetRate(v float32) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetRect(v *DetectLivingFaceResponseBodyDataElementsResultsRect) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Rect = v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetSuggestion(v string) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

type DetectLivingFaceResponseBodyDataElementsResultsFrames struct {
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Url  *string  `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElementsResultsFrames) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElementsResultsFrames) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsFrames) SetRate(v float32) *DetectLivingFaceResponseBodyDataElementsResultsFrames {
	s.Rate = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsFrames) SetUrl(v string) *DetectLivingFaceResponseBodyDataElementsResultsFrames {
	s.Url = &v
	return s
}

type DetectLivingFaceResponseBodyDataElementsResultsRect struct {
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElementsResultsRect) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElementsResultsRect) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) SetHeight(v int64) *DetectLivingFaceResponseBodyDataElementsResultsRect {
	s.Height = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) SetLeft(v int64) *DetectLivingFaceResponseBodyDataElementsResultsRect {
	s.Left = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) SetTop(v int64) *DetectLivingFaceResponseBodyDataElementsResultsRect {
	s.Top = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) SetWidth(v int64) *DetectLivingFaceResponseBodyDataElementsResultsRect {
	s.Width = &v
	return s
}

type DetectLivingFaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectLivingFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectLivingFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectLivingFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponse) SetHeaders(v map[string]*string) *DetectLivingFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectLivingFaceResponse) SetStatusCode(v int32) *DetectLivingFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectLivingFaceResponse) SetBody(v *DetectLivingFaceResponseBody) *DetectLivingFaceResponse {
	s.Body = v
	return s
}

type DetectPedestrianRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectPedestrianRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianRequest) SetImageURL(v string) *DetectPedestrianRequest {
	s.ImageURL = &v
	return s
}

type DetectPedestrianAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectPedestrianAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianAdvanceRequest) SetImageURLObject(v io.Reader) *DetectPedestrianAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectPedestrianResponseBody struct {
	Data      *DetectPedestrianResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectPedestrianResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianResponseBody) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponseBody) SetData(v *DetectPedestrianResponseBodyData) *DetectPedestrianResponseBody {
	s.Data = v
	return s
}

func (s *DetectPedestrianResponseBody) SetRequestId(v string) *DetectPedestrianResponseBody {
	s.RequestId = &v
	return s
}

type DetectPedestrianResponseBodyData struct {
	Elements []*DetectPedestrianResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Height   *int32                                      `json:"Height,omitempty" xml:"Height,omitempty"`
	Width    *int32                                      `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectPedestrianResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponseBodyData) SetElements(v []*DetectPedestrianResponseBodyDataElements) *DetectPedestrianResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectPedestrianResponseBodyData) SetHeight(v int32) *DetectPedestrianResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectPedestrianResponseBodyData) SetWidth(v int32) *DetectPedestrianResponseBodyData {
	s.Width = &v
	return s
}

type DetectPedestrianResponseBodyDataElements struct {
	// 1
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type  *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectPedestrianResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponseBodyDataElements) SetBoxes(v []*int32) *DetectPedestrianResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectPedestrianResponseBodyDataElements) SetScore(v float32) *DetectPedestrianResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectPedestrianResponseBodyDataElements) SetType(v string) *DetectPedestrianResponseBodyDataElements {
	s.Type = &v
	return s
}

type DetectPedestrianResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectPedestrianResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectPedestrianResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianResponse) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponse) SetHeaders(v map[string]*string) *DetectPedestrianResponse {
	s.Headers = v
	return s
}

func (s *DetectPedestrianResponse) SetStatusCode(v int32) *DetectPedestrianResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectPedestrianResponse) SetBody(v *DetectPedestrianResponseBody) *DetectPedestrianResponse {
	s.Body = v
	return s
}

type DetectPedestrianIntrusionRequest struct {
	DetectRegion []*DetectPedestrianIntrusionRequestDetectRegion `json:"DetectRegion,omitempty" xml:"DetectRegion,omitempty" type:"Repeated"`
	ImageURL     *string                                         `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RegionType   *string                                         `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
}

func (s DetectPedestrianIntrusionRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionRequest) SetDetectRegion(v []*DetectPedestrianIntrusionRequestDetectRegion) *DetectPedestrianIntrusionRequest {
	s.DetectRegion = v
	return s
}

func (s *DetectPedestrianIntrusionRequest) SetImageURL(v string) *DetectPedestrianIntrusionRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectPedestrianIntrusionRequest) SetRegionType(v string) *DetectPedestrianIntrusionRequest {
	s.RegionType = &v
	return s
}

type DetectPedestrianIntrusionRequestDetectRegion struct {
	Line *DetectPedestrianIntrusionRequestDetectRegionLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	Rect *DetectPedestrianIntrusionRequestDetectRegionRect `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
}

func (s DetectPedestrianIntrusionRequestDetectRegion) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionRequestDetectRegion) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionRequestDetectRegion) SetLine(v *DetectPedestrianIntrusionRequestDetectRegionLine) *DetectPedestrianIntrusionRequestDetectRegion {
	s.Line = v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegion) SetRect(v *DetectPedestrianIntrusionRequestDetectRegionRect) *DetectPedestrianIntrusionRequestDetectRegion {
	s.Rect = v
	return s
}

type DetectPedestrianIntrusionRequestDetectRegionLine struct {
	X1 *int64 `json:"X1,omitempty" xml:"X1,omitempty"`
	X2 *int64 `json:"X2,omitempty" xml:"X2,omitempty"`
	Y1 *int64 `json:"Y1,omitempty" xml:"Y1,omitempty"`
	Y2 *int64 `json:"Y2,omitempty" xml:"Y2,omitempty"`
}

func (s DetectPedestrianIntrusionRequestDetectRegionLine) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionRequestDetectRegionLine) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionRequestDetectRegionLine) SetX1(v int64) *DetectPedestrianIntrusionRequestDetectRegionLine {
	s.X1 = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionLine) SetX2(v int64) *DetectPedestrianIntrusionRequestDetectRegionLine {
	s.X2 = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionLine) SetY1(v int64) *DetectPedestrianIntrusionRequestDetectRegionLine {
	s.Y1 = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionLine) SetY2(v int64) *DetectPedestrianIntrusionRequestDetectRegionLine {
	s.Y2 = &v
	return s
}

type DetectPedestrianIntrusionRequestDetectRegionRect struct {
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DetectPedestrianIntrusionRequestDetectRegionRect) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionRequestDetectRegionRect) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionRequestDetectRegionRect) SetBottom(v int64) *DetectPedestrianIntrusionRequestDetectRegionRect {
	s.Bottom = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionRect) SetLeft(v int64) *DetectPedestrianIntrusionRequestDetectRegionRect {
	s.Left = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionRect) SetRight(v int64) *DetectPedestrianIntrusionRequestDetectRegionRect {
	s.Right = &v
	return s
}

func (s *DetectPedestrianIntrusionRequestDetectRegionRect) SetTop(v int64) *DetectPedestrianIntrusionRequestDetectRegionRect {
	s.Top = &v
	return s
}

type DetectPedestrianIntrusionAdvanceRequest struct {
	DetectRegion   []*DetectPedestrianIntrusionAdvanceRequestDetectRegion `json:"DetectRegion,omitempty" xml:"DetectRegion,omitempty" type:"Repeated"`
	ImageURLObject io.Reader                                              `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RegionType     *string                                                `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
}

func (s DetectPedestrianIntrusionAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionAdvanceRequest) SetDetectRegion(v []*DetectPedestrianIntrusionAdvanceRequestDetectRegion) *DetectPedestrianIntrusionAdvanceRequest {
	s.DetectRegion = v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequest) SetImageURLObject(v io.Reader) *DetectPedestrianIntrusionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequest) SetRegionType(v string) *DetectPedestrianIntrusionAdvanceRequest {
	s.RegionType = &v
	return s
}

type DetectPedestrianIntrusionAdvanceRequestDetectRegion struct {
	Line *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	Rect *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegion) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegion) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegion) SetLine(v *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) *DetectPedestrianIntrusionAdvanceRequestDetectRegion {
	s.Line = v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegion) SetRect(v *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) *DetectPedestrianIntrusionAdvanceRequestDetectRegion {
	s.Rect = v
	return s
}

type DetectPedestrianIntrusionAdvanceRequestDetectRegionLine struct {
	X1 *int64 `json:"X1,omitempty" xml:"X1,omitempty"`
	X2 *int64 `json:"X2,omitempty" xml:"X2,omitempty"`
	Y1 *int64 `json:"Y1,omitempty" xml:"Y1,omitempty"`
	Y2 *int64 `json:"Y2,omitempty" xml:"Y2,omitempty"`
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) SetX1(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine {
	s.X1 = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) SetX2(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine {
	s.X2 = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) SetY1(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine {
	s.Y1 = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine) SetY2(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionLine {
	s.Y2 = &v
	return s
}

type DetectPedestrianIntrusionAdvanceRequestDetectRegionRect struct {
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) SetBottom(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect {
	s.Bottom = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) SetLeft(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect {
	s.Left = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) SetRight(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect {
	s.Right = &v
	return s
}

func (s *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect) SetTop(v int64) *DetectPedestrianIntrusionAdvanceRequestDetectRegionRect {
	s.Top = &v
	return s
}

type DetectPedestrianIntrusionShrinkRequest struct {
	DetectRegionShrink *string `json:"DetectRegion,omitempty" xml:"DetectRegion,omitempty"`
	ImageURL           *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RegionType         *string `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
}

func (s DetectPedestrianIntrusionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionShrinkRequest) SetDetectRegionShrink(v string) *DetectPedestrianIntrusionShrinkRequest {
	s.DetectRegionShrink = &v
	return s
}

func (s *DetectPedestrianIntrusionShrinkRequest) SetImageURL(v string) *DetectPedestrianIntrusionShrinkRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectPedestrianIntrusionShrinkRequest) SetRegionType(v string) *DetectPedestrianIntrusionShrinkRequest {
	s.RegionType = &v
	return s
}

type DetectPedestrianIntrusionResponseBody struct {
	Data      *DetectPedestrianIntrusionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectPedestrianIntrusionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponseBody) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponseBody) SetData(v *DetectPedestrianIntrusionResponseBodyData) *DetectPedestrianIntrusionResponseBody {
	s.Data = v
	return s
}

func (s *DetectPedestrianIntrusionResponseBody) SetRequestId(v string) *DetectPedestrianIntrusionResponseBody {
	s.RequestId = &v
	return s
}

type DetectPedestrianIntrusionResponseBodyData struct {
	Elements    []*DetectPedestrianIntrusionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	ImageHeight *int64                                               `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageWidth  *int64                                               `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
}

func (s DetectPedestrianIntrusionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponseBodyData) SetElements(v []*DetectPedestrianIntrusionResponseBodyDataElements) *DetectPedestrianIntrusionResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyData) SetImageHeight(v int64) *DetectPedestrianIntrusionResponseBodyData {
	s.ImageHeight = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyData) SetImageWidth(v int64) *DetectPedestrianIntrusionResponseBodyData {
	s.ImageWidth = &v
	return s
}

type DetectPedestrianIntrusionResponseBodyDataElements struct {
	Box       *DetectPedestrianIntrusionResponseBodyDataElementsBox `json:"Box,omitempty" xml:"Box,omitempty" type:"Struct"`
	BoxId     *int64                                                `json:"BoxId,omitempty" xml:"BoxId,omitempty"`
	IsIntrude *bool                                                 `json:"IsIntrude,omitempty" xml:"IsIntrude,omitempty"`
	Score     *int64                                                `json:"Score,omitempty" xml:"Score,omitempty"`
	Type      *string                                               `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectPedestrianIntrusionResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetBox(v *DetectPedestrianIntrusionResponseBodyDataElementsBox) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.Box = v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetBoxId(v int64) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.BoxId = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetIsIntrude(v bool) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.IsIntrude = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetScore(v int64) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElements) SetType(v string) *DetectPedestrianIntrusionResponseBodyDataElements {
	s.Type = &v
	return s
}

type DetectPedestrianIntrusionResponseBodyDataElementsBox struct {
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Right  *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DetectPedestrianIntrusionResponseBodyDataElementsBox) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponseBodyDataElementsBox) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponseBodyDataElementsBox) SetBottom(v int64) *DetectPedestrianIntrusionResponseBodyDataElementsBox {
	s.Bottom = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElementsBox) SetLeft(v int64) *DetectPedestrianIntrusionResponseBodyDataElementsBox {
	s.Left = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElementsBox) SetRight(v int64) *DetectPedestrianIntrusionResponseBodyDataElementsBox {
	s.Right = &v
	return s
}

func (s *DetectPedestrianIntrusionResponseBodyDataElementsBox) SetTop(v int64) *DetectPedestrianIntrusionResponseBodyDataElementsBox {
	s.Top = &v
	return s
}

type DetectPedestrianIntrusionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectPedestrianIntrusionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectPedestrianIntrusionResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectPedestrianIntrusionResponse) GoString() string {
	return s.String()
}

func (s *DetectPedestrianIntrusionResponse) SetHeaders(v map[string]*string) *DetectPedestrianIntrusionResponse {
	s.Headers = v
	return s
}

func (s *DetectPedestrianIntrusionResponse) SetStatusCode(v int32) *DetectPedestrianIntrusionResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectPedestrianIntrusionResponse) SetBody(v *DetectPedestrianIntrusionResponseBody) *DetectPedestrianIntrusionResponse {
	s.Body = v
	return s
}

type DetectVideoLivingFaceRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s DetectVideoLivingFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceRequest) SetVideoUrl(v string) *DetectVideoLivingFaceRequest {
	s.VideoUrl = &v
	return s
}

type DetectVideoLivingFaceAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s DetectVideoLivingFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceAdvanceRequest) SetVideoUrlObject(v io.Reader) *DetectVideoLivingFaceAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type DetectVideoLivingFaceResponseBody struct {
	Data      *DetectVideoLivingFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVideoLivingFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponseBody) SetData(v *DetectVideoLivingFaceResponseBodyData) *DetectVideoLivingFaceResponseBody {
	s.Data = v
	return s
}

func (s *DetectVideoLivingFaceResponseBody) SetRequestId(v string) *DetectVideoLivingFaceResponseBody {
	s.RequestId = &v
	return s
}

type DetectVideoLivingFaceResponseBodyData struct {
	Elements []*DetectVideoLivingFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectVideoLivingFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponseBodyData) SetElements(v []*DetectVideoLivingFaceResponseBodyDataElements) *DetectVideoLivingFaceResponseBodyData {
	s.Elements = v
	return s
}

type DetectVideoLivingFaceResponseBodyDataElements struct {
	FaceConfidence *float32 `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	LiveConfidence *float32 `json:"LiveConfidence,omitempty" xml:"LiveConfidence,omitempty"`
	Rect           []*int32 `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Repeated"`
}

func (s DetectVideoLivingFaceResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) SetFaceConfidence(v float32) *DetectVideoLivingFaceResponseBodyDataElements {
	s.FaceConfidence = &v
	return s
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) SetLiveConfidence(v float32) *DetectVideoLivingFaceResponseBodyDataElements {
	s.LiveConfidence = &v
	return s
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) SetRect(v []*int32) *DetectVideoLivingFaceResponseBodyDataElements {
	s.Rect = v
	return s
}

type DetectVideoLivingFaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectVideoLivingFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectVideoLivingFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoLivingFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponse) SetHeaders(v map[string]*string) *DetectVideoLivingFaceResponse {
	s.Headers = v
	return s
}

func (s *DetectVideoLivingFaceResponse) SetStatusCode(v int32) *DetectVideoLivingFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVideoLivingFaceResponse) SetBody(v *DetectVideoLivingFaceResponseBody) *DetectVideoLivingFaceResponse {
	s.Body = v
	return s
}

type EnhanceFaceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceRequest) GoString() string {
	return s.String()
}

func (s *EnhanceFaceRequest) SetImageURL(v string) *EnhanceFaceRequest {
	s.ImageURL = &v
	return s
}

type EnhanceFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EnhanceFaceAdvanceRequest) SetImageURLObject(v io.Reader) *EnhanceFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type EnhanceFaceResponseBody struct {
	Data      *EnhanceFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnhanceFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceResponseBody) GoString() string {
	return s.String()
}

func (s *EnhanceFaceResponseBody) SetData(v *EnhanceFaceResponseBodyData) *EnhanceFaceResponseBody {
	s.Data = v
	return s
}

func (s *EnhanceFaceResponseBody) SetRequestId(v string) *EnhanceFaceResponseBody {
	s.RequestId = &v
	return s
}

type EnhanceFaceResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnhanceFaceResponseBodyData) SetImageURL(v string) *EnhanceFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

type EnhanceFaceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnhanceFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnhanceFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnhanceFaceResponse) GoString() string {
	return s.String()
}

func (s *EnhanceFaceResponse) SetHeaders(v map[string]*string) *EnhanceFaceResponse {
	s.Headers = v
	return s
}

func (s *EnhanceFaceResponse) SetStatusCode(v int32) *EnhanceFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *EnhanceFaceResponse) SetBody(v *EnhanceFaceResponseBody) *EnhanceFaceResponse {
	s.Body = v
	return s
}

type ExtractFingerPrintRequest struct {
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ExtractFingerPrintRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractFingerPrintRequest) GoString() string {
	return s.String()
}

func (s *ExtractFingerPrintRequest) SetImageData(v string) *ExtractFingerPrintRequest {
	s.ImageData = &v
	return s
}

func (s *ExtractFingerPrintRequest) SetImageURL(v string) *ExtractFingerPrintRequest {
	s.ImageURL = &v
	return s
}

type ExtractFingerPrintAdvanceRequest struct {
	ImageData      *string   `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ExtractFingerPrintAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractFingerPrintAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ExtractFingerPrintAdvanceRequest) SetImageData(v string) *ExtractFingerPrintAdvanceRequest {
	s.ImageData = &v
	return s
}

func (s *ExtractFingerPrintAdvanceRequest) SetImageURLObject(v io.Reader) *ExtractFingerPrintAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type ExtractFingerPrintResponseBody struct {
	Data      *ExtractFingerPrintResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtractFingerPrintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractFingerPrintResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractFingerPrintResponseBody) SetData(v *ExtractFingerPrintResponseBodyData) *ExtractFingerPrintResponseBody {
	s.Data = v
	return s
}

func (s *ExtractFingerPrintResponseBody) SetRequestId(v string) *ExtractFingerPrintResponseBody {
	s.RequestId = &v
	return s
}

type ExtractFingerPrintResponseBodyData struct {
	FingerPrint []byte `json:"FingerPrint,omitempty" xml:"FingerPrint,omitempty"`
}

func (s ExtractFingerPrintResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExtractFingerPrintResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExtractFingerPrintResponseBodyData) SetFingerPrint(v []byte) *ExtractFingerPrintResponseBodyData {
	s.FingerPrint = v
	return s
}

type ExtractFingerPrintResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExtractFingerPrintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractFingerPrintResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractFingerPrintResponse) GoString() string {
	return s.String()
}

func (s *ExtractFingerPrintResponse) SetHeaders(v map[string]*string) *ExtractFingerPrintResponse {
	s.Headers = v
	return s
}

func (s *ExtractFingerPrintResponse) SetStatusCode(v int32) *ExtractFingerPrintResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtractFingerPrintResponse) SetBody(v *ExtractFingerPrintResponseBody) *ExtractFingerPrintResponse {
	s.Body = v
	return s
}

type ExtractPedestrianFeatureAttrRequest struct {
	ImageURL       *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Mode           *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ExtractPedestrianFeatureAttrRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrRequest) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrRequest) SetImageURL(v string) *ExtractPedestrianFeatureAttrRequest {
	s.ImageURL = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrRequest) SetMode(v string) *ExtractPedestrianFeatureAttrRequest {
	s.Mode = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrRequest) SetServiceVersion(v string) *ExtractPedestrianFeatureAttrRequest {
	s.ServiceVersion = &v
	return s
}

type ExtractPedestrianFeatureAttrAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ServiceVersion *string   `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ExtractPedestrianFeatureAttrAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrAdvanceRequest) SetImageURLObject(v io.Reader) *ExtractPedestrianFeatureAttrAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *ExtractPedestrianFeatureAttrAdvanceRequest) SetMode(v string) *ExtractPedestrianFeatureAttrAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrAdvanceRequest) SetServiceVersion(v string) *ExtractPedestrianFeatureAttrAdvanceRequest {
	s.ServiceVersion = &v
	return s
}

type ExtractPedestrianFeatureAttrResponseBody struct {
	Data      *ExtractPedestrianFeatureAttrResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtractPedestrianFeatureAttrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrResponseBody) SetData(v *ExtractPedestrianFeatureAttrResponseBodyData) *ExtractPedestrianFeatureAttrResponseBody {
	s.Data = v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBody) SetRequestId(v string) *ExtractPedestrianFeatureAttrResponseBody {
	s.RequestId = &v
	return s
}

type ExtractPedestrianFeatureAttrResponseBodyData struct {
	Age              *string  `json:"Age,omitempty" xml:"Age,omitempty"`
	AgeScore         *float32 `json:"AgeScore,omitempty" xml:"AgeScore,omitempty"`
	Feature          *string  `json:"Feature,omitempty" xml:"Feature,omitempty"`
	Gender           *string  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	GenderScore      *float32 `json:"GenderScore,omitempty" xml:"GenderScore,omitempty"`
	Hair             *string  `json:"Hair,omitempty" xml:"Hair,omitempty"`
	HairScore        *float32 `json:"HairScore,omitempty" xml:"HairScore,omitempty"`
	LowerColor       *string  `json:"LowerColor,omitempty" xml:"LowerColor,omitempty"`
	LowerColorScore  *float32 `json:"LowerColorScore,omitempty" xml:"LowerColorScore,omitempty"`
	LowerType        *string  `json:"LowerType,omitempty" xml:"LowerType,omitempty"`
	LowerTypeScore   *float32 `json:"LowerTypeScore,omitempty" xml:"LowerTypeScore,omitempty"`
	ObjType          *string  `json:"ObjType,omitempty" xml:"ObjType,omitempty"`
	ObjTypeScore     *float32 `json:"ObjTypeScore,omitempty" xml:"ObjTypeScore,omitempty"`
	Orientation      *string  `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	OrientationScore *float32 `json:"OrientationScore,omitempty" xml:"OrientationScore,omitempty"`
	QualityScore     *float32 `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	UpperColor       *string  `json:"UpperColor,omitempty" xml:"UpperColor,omitempty"`
	UpperColorScore  *float32 `json:"UpperColorScore,omitempty" xml:"UpperColorScore,omitempty"`
	UpperType        *string  `json:"UpperType,omitempty" xml:"UpperType,omitempty"`
	UpperTypeScore   *float32 `json:"UpperTypeScore,omitempty" xml:"UpperTypeScore,omitempty"`
}

func (s ExtractPedestrianFeatureAttrResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetAge(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.Age = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetAgeScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.AgeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetFeature(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.Feature = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetGender(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.Gender = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetGenderScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.GenderScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetHair(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.Hair = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetHairScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.HairScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetLowerColor(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.LowerColor = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetLowerColorScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.LowerColorScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetLowerType(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.LowerType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetLowerTypeScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.LowerTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetObjType(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.ObjType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetObjTypeScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.ObjTypeScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetOrientation(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.Orientation = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetOrientationScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.OrientationScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetQualityScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.QualityScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetUpperColor(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.UpperColor = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetUpperColorScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.UpperColorScore = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetUpperType(v string) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.UpperType = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponseBodyData) SetUpperTypeScore(v float32) *ExtractPedestrianFeatureAttrResponseBodyData {
	s.UpperTypeScore = &v
	return s
}

type ExtractPedestrianFeatureAttrResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExtractPedestrianFeatureAttrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractPedestrianFeatureAttrResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractPedestrianFeatureAttrResponse) GoString() string {
	return s.String()
}

func (s *ExtractPedestrianFeatureAttrResponse) SetHeaders(v map[string]*string) *ExtractPedestrianFeatureAttrResponse {
	s.Headers = v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponse) SetStatusCode(v int32) *ExtractPedestrianFeatureAttrResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtractPedestrianFeatureAttrResponse) SetBody(v *ExtractPedestrianFeatureAttrResponseBody) *ExtractPedestrianFeatureAttrResponse {
	s.Body = v
	return s
}

type FaceBeautyRequest struct {
	ImageURL *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Sharp    *float32 `json:"Sharp,omitempty" xml:"Sharp,omitempty"`
	Smooth   *float32 `json:"Smooth,omitempty" xml:"Smooth,omitempty"`
	White    *float32 `json:"White,omitempty" xml:"White,omitempty"`
}

func (s FaceBeautyRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyRequest) GoString() string {
	return s.String()
}

func (s *FaceBeautyRequest) SetImageURL(v string) *FaceBeautyRequest {
	s.ImageURL = &v
	return s
}

func (s *FaceBeautyRequest) SetSharp(v float32) *FaceBeautyRequest {
	s.Sharp = &v
	return s
}

func (s *FaceBeautyRequest) SetSmooth(v float32) *FaceBeautyRequest {
	s.Smooth = &v
	return s
}

func (s *FaceBeautyRequest) SetWhite(v float32) *FaceBeautyRequest {
	s.White = &v
	return s
}

type FaceBeautyAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Sharp          *float32  `json:"Sharp,omitempty" xml:"Sharp,omitempty"`
	Smooth         *float32  `json:"Smooth,omitempty" xml:"Smooth,omitempty"`
	White          *float32  `json:"White,omitempty" xml:"White,omitempty"`
}

func (s FaceBeautyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceBeautyAdvanceRequest) SetImageURLObject(v io.Reader) *FaceBeautyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *FaceBeautyAdvanceRequest) SetSharp(v float32) *FaceBeautyAdvanceRequest {
	s.Sharp = &v
	return s
}

func (s *FaceBeautyAdvanceRequest) SetSmooth(v float32) *FaceBeautyAdvanceRequest {
	s.Smooth = &v
	return s
}

func (s *FaceBeautyAdvanceRequest) SetWhite(v float32) *FaceBeautyAdvanceRequest {
	s.White = &v
	return s
}

type FaceBeautyResponseBody struct {
	Data      *FaceBeautyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FaceBeautyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyResponseBody) GoString() string {
	return s.String()
}

func (s *FaceBeautyResponseBody) SetData(v *FaceBeautyResponseBodyData) *FaceBeautyResponseBody {
	s.Data = v
	return s
}

func (s *FaceBeautyResponseBody) SetRequestId(v string) *FaceBeautyResponseBody {
	s.RequestId = &v
	return s
}

type FaceBeautyResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s FaceBeautyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceBeautyResponseBodyData) SetImageURL(v string) *FaceBeautyResponseBodyData {
	s.ImageURL = &v
	return s
}

type FaceBeautyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FaceBeautyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceBeautyResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceBeautyResponse) GoString() string {
	return s.String()
}

func (s *FaceBeautyResponse) SetHeaders(v map[string]*string) *FaceBeautyResponse {
	s.Headers = v
	return s
}

func (s *FaceBeautyResponse) SetStatusCode(v int32) *FaceBeautyResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceBeautyResponse) SetBody(v *FaceBeautyResponseBody) *FaceBeautyResponse {
	s.Body = v
	return s
}

type FaceFilterRequest struct {
	ImageURL     *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ResourceType *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Strength     *float32 `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterRequest) GoString() string {
	return s.String()
}

func (s *FaceFilterRequest) SetImageURL(v string) *FaceFilterRequest {
	s.ImageURL = &v
	return s
}

func (s *FaceFilterRequest) SetResourceType(v string) *FaceFilterRequest {
	s.ResourceType = &v
	return s
}

func (s *FaceFilterRequest) SetStrength(v float32) *FaceFilterRequest {
	s.Strength = &v
	return s
}

type FaceFilterAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ResourceType   *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Strength       *float32  `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceFilterAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceFilterAdvanceRequest) SetImageURLObject(v io.Reader) *FaceFilterAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *FaceFilterAdvanceRequest) SetResourceType(v string) *FaceFilterAdvanceRequest {
	s.ResourceType = &v
	return s
}

func (s *FaceFilterAdvanceRequest) SetStrength(v float32) *FaceFilterAdvanceRequest {
	s.Strength = &v
	return s
}

type FaceFilterResponseBody struct {
	Data      *FaceFilterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FaceFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterResponseBody) GoString() string {
	return s.String()
}

func (s *FaceFilterResponseBody) SetData(v *FaceFilterResponseBodyData) *FaceFilterResponseBody {
	s.Data = v
	return s
}

func (s *FaceFilterResponseBody) SetRequestId(v string) *FaceFilterResponseBody {
	s.RequestId = &v
	return s
}

type FaceFilterResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s FaceFilterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceFilterResponseBodyData) SetImageURL(v string) *FaceFilterResponseBodyData {
	s.ImageURL = &v
	return s
}

type FaceFilterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FaceFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceFilterResponse) GoString() string {
	return s.String()
}

func (s *FaceFilterResponse) SetHeaders(v map[string]*string) *FaceFilterResponse {
	s.Headers = v
	return s
}

func (s *FaceFilterResponse) SetStatusCode(v int32) *FaceFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceFilterResponse) SetBody(v *FaceFilterResponseBody) *FaceFilterResponse {
	s.Body = v
	return s
}

type FaceMakeupRequest struct {
	ImageURL     *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	MakeupType   *string  `json:"MakeupType,omitempty" xml:"MakeupType,omitempty"`
	ResourceType *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Strength     *float32 `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceMakeupRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupRequest) GoString() string {
	return s.String()
}

func (s *FaceMakeupRequest) SetImageURL(v string) *FaceMakeupRequest {
	s.ImageURL = &v
	return s
}

func (s *FaceMakeupRequest) SetMakeupType(v string) *FaceMakeupRequest {
	s.MakeupType = &v
	return s
}

func (s *FaceMakeupRequest) SetResourceType(v string) *FaceMakeupRequest {
	s.ResourceType = &v
	return s
}

func (s *FaceMakeupRequest) SetStrength(v float32) *FaceMakeupRequest {
	s.Strength = &v
	return s
}

type FaceMakeupAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	MakeupType     *string   `json:"MakeupType,omitempty" xml:"MakeupType,omitempty"`
	ResourceType   *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Strength       *float32  `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceMakeupAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceMakeupAdvanceRequest) SetImageURLObject(v io.Reader) *FaceMakeupAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *FaceMakeupAdvanceRequest) SetMakeupType(v string) *FaceMakeupAdvanceRequest {
	s.MakeupType = &v
	return s
}

func (s *FaceMakeupAdvanceRequest) SetResourceType(v string) *FaceMakeupAdvanceRequest {
	s.ResourceType = &v
	return s
}

func (s *FaceMakeupAdvanceRequest) SetStrength(v float32) *FaceMakeupAdvanceRequest {
	s.Strength = &v
	return s
}

type FaceMakeupResponseBody struct {
	Data      *FaceMakeupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FaceMakeupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupResponseBody) GoString() string {
	return s.String()
}

func (s *FaceMakeupResponseBody) SetData(v *FaceMakeupResponseBodyData) *FaceMakeupResponseBody {
	s.Data = v
	return s
}

func (s *FaceMakeupResponseBody) SetRequestId(v string) *FaceMakeupResponseBody {
	s.RequestId = &v
	return s
}

type FaceMakeupResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s FaceMakeupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceMakeupResponseBodyData) SetImageURL(v string) *FaceMakeupResponseBodyData {
	s.ImageURL = &v
	return s
}

type FaceMakeupResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FaceMakeupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceMakeupResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceMakeupResponse) GoString() string {
	return s.String()
}

func (s *FaceMakeupResponse) SetHeaders(v map[string]*string) *FaceMakeupResponse {
	s.Headers = v
	return s
}

func (s *FaceMakeupResponse) SetStatusCode(v int32) *FaceMakeupResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceMakeupResponse) SetBody(v *FaceMakeupResponseBody) *FaceMakeupResponse {
	s.Body = v
	return s
}

type FaceTidyupRequest struct {
	ImageURL  *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ShapeType *int32   `json:"ShapeType,omitempty" xml:"ShapeType,omitempty"`
	Strength  *float32 `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceTidyupRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupRequest) GoString() string {
	return s.String()
}

func (s *FaceTidyupRequest) SetImageURL(v string) *FaceTidyupRequest {
	s.ImageURL = &v
	return s
}

func (s *FaceTidyupRequest) SetShapeType(v int32) *FaceTidyupRequest {
	s.ShapeType = &v
	return s
}

func (s *FaceTidyupRequest) SetStrength(v float32) *FaceTidyupRequest {
	s.Strength = &v
	return s
}

type FaceTidyupAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ShapeType      *int32    `json:"ShapeType,omitempty" xml:"ShapeType,omitempty"`
	Strength       *float32  `json:"Strength,omitempty" xml:"Strength,omitempty"`
}

func (s FaceTidyupAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceTidyupAdvanceRequest) SetImageURLObject(v io.Reader) *FaceTidyupAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *FaceTidyupAdvanceRequest) SetShapeType(v int32) *FaceTidyupAdvanceRequest {
	s.ShapeType = &v
	return s
}

func (s *FaceTidyupAdvanceRequest) SetStrength(v float32) *FaceTidyupAdvanceRequest {
	s.Strength = &v
	return s
}

type FaceTidyupResponseBody struct {
	Data      *FaceTidyupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FaceTidyupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupResponseBody) GoString() string {
	return s.String()
}

func (s *FaceTidyupResponseBody) SetData(v *FaceTidyupResponseBodyData) *FaceTidyupResponseBody {
	s.Data = v
	return s
}

func (s *FaceTidyupResponseBody) SetRequestId(v string) *FaceTidyupResponseBody {
	s.RequestId = &v
	return s
}

type FaceTidyupResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s FaceTidyupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceTidyupResponseBodyData) SetImageURL(v string) *FaceTidyupResponseBodyData {
	s.ImageURL = &v
	return s
}

type FaceTidyupResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FaceTidyupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceTidyupResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceTidyupResponse) GoString() string {
	return s.String()
}

func (s *FaceTidyupResponse) SetHeaders(v map[string]*string) *FaceTidyupResponse {
	s.Headers = v
	return s
}

func (s *FaceTidyupResponse) SetStatusCode(v int32) *FaceTidyupResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceTidyupResponse) SetBody(v *FaceTidyupResponseBody) *FaceTidyupResponse {
	s.Body = v
	return s
}

type GenRealPersonVerificationTokenRequest struct {
	CertificateName   *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	CertificateNumber *string `json:"CertificateNumber,omitempty" xml:"CertificateNumber,omitempty"`
	MetaInfo          *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
}

func (s GenRealPersonVerificationTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GenRealPersonVerificationTokenRequest) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenRequest) SetCertificateName(v string) *GenRealPersonVerificationTokenRequest {
	s.CertificateName = &v
	return s
}

func (s *GenRealPersonVerificationTokenRequest) SetCertificateNumber(v string) *GenRealPersonVerificationTokenRequest {
	s.CertificateNumber = &v
	return s
}

func (s *GenRealPersonVerificationTokenRequest) SetMetaInfo(v string) *GenRealPersonVerificationTokenRequest {
	s.MetaInfo = &v
	return s
}

type GenRealPersonVerificationTokenResponseBody struct {
	Data      *GenRealPersonVerificationTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenRealPersonVerificationTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenRealPersonVerificationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenResponseBody) SetData(v *GenRealPersonVerificationTokenResponseBodyData) *GenRealPersonVerificationTokenResponseBody {
	s.Data = v
	return s
}

func (s *GenRealPersonVerificationTokenResponseBody) SetRequestId(v string) *GenRealPersonVerificationTokenResponseBody {
	s.RequestId = &v
	return s
}

type GenRealPersonVerificationTokenResponseBodyData struct {
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s GenRealPersonVerificationTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenRealPersonVerificationTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenResponseBodyData) SetVerificationToken(v string) *GenRealPersonVerificationTokenResponseBodyData {
	s.VerificationToken = &v
	return s
}

type GenRealPersonVerificationTokenResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenRealPersonVerificationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenRealPersonVerificationTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GenRealPersonVerificationTokenResponse) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenResponse) SetHeaders(v map[string]*string) *GenRealPersonVerificationTokenResponse {
	s.Headers = v
	return s
}

func (s *GenRealPersonVerificationTokenResponse) SetStatusCode(v int32) *GenRealPersonVerificationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenRealPersonVerificationTokenResponse) SetBody(v *GenRealPersonVerificationTokenResponseBody) *GenRealPersonVerificationTokenResponse {
	s.Body = v
	return s
}

type GenerateHumanAnimeStyleRequest struct {
	AlgoType *string `json:"AlgoType,omitempty" xml:"AlgoType,omitempty"`
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanAnimeStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleRequest) SetAlgoType(v string) *GenerateHumanAnimeStyleRequest {
	s.AlgoType = &v
	return s
}

func (s *GenerateHumanAnimeStyleRequest) SetImageURL(v string) *GenerateHumanAnimeStyleRequest {
	s.ImageURL = &v
	return s
}

type GenerateHumanAnimeStyleAdvanceRequest struct {
	AlgoType       *string   `json:"AlgoType,omitempty" xml:"AlgoType,omitempty"`
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanAnimeStyleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleAdvanceRequest) SetAlgoType(v string) *GenerateHumanAnimeStyleAdvanceRequest {
	s.AlgoType = &v
	return s
}

func (s *GenerateHumanAnimeStyleAdvanceRequest) SetImageURLObject(v io.Reader) *GenerateHumanAnimeStyleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type GenerateHumanAnimeStyleResponseBody struct {
	Data      *GenerateHumanAnimeStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateHumanAnimeStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleResponseBody) SetData(v *GenerateHumanAnimeStyleResponseBodyData) *GenerateHumanAnimeStyleResponseBody {
	s.Data = v
	return s
}

func (s *GenerateHumanAnimeStyleResponseBody) SetRequestId(v string) *GenerateHumanAnimeStyleResponseBody {
	s.RequestId = &v
	return s
}

type GenerateHumanAnimeStyleResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanAnimeStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleResponseBodyData) SetImageURL(v string) *GenerateHumanAnimeStyleResponseBodyData {
	s.ImageURL = &v
	return s
}

type GenerateHumanAnimeStyleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateHumanAnimeStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateHumanAnimeStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanAnimeStyleResponse) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleResponse) SetHeaders(v map[string]*string) *GenerateHumanAnimeStyleResponse {
	s.Headers = v
	return s
}

func (s *GenerateHumanAnimeStyleResponse) SetStatusCode(v int32) *GenerateHumanAnimeStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateHumanAnimeStyleResponse) SetBody(v *GenerateHumanAnimeStyleResponseBody) *GenerateHumanAnimeStyleResponse {
	s.Body = v
	return s
}

type GenerateHumanSketchStyleRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReturnType *string `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
}

func (s GenerateHumanSketchStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleRequest) SetImageURL(v string) *GenerateHumanSketchStyleRequest {
	s.ImageURL = &v
	return s
}

func (s *GenerateHumanSketchStyleRequest) SetReturnType(v string) *GenerateHumanSketchStyleRequest {
	s.ReturnType = &v
	return s
}

type GenerateHumanSketchStyleAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	ReturnType     *string   `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
}

func (s GenerateHumanSketchStyleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleAdvanceRequest) SetImageURLObject(v io.Reader) *GenerateHumanSketchStyleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *GenerateHumanSketchStyleAdvanceRequest) SetReturnType(v string) *GenerateHumanSketchStyleAdvanceRequest {
	s.ReturnType = &v
	return s
}

type GenerateHumanSketchStyleResponseBody struct {
	Data      *GenerateHumanSketchStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateHumanSketchStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleResponseBody) SetData(v *GenerateHumanSketchStyleResponseBodyData) *GenerateHumanSketchStyleResponseBody {
	s.Data = v
	return s
}

func (s *GenerateHumanSketchStyleResponseBody) SetRequestId(v string) *GenerateHumanSketchStyleResponseBody {
	s.RequestId = &v
	return s
}

type GenerateHumanSketchStyleResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanSketchStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleResponseBodyData) SetImageURL(v string) *GenerateHumanSketchStyleResponseBodyData {
	s.ImageURL = &v
	return s
}

type GenerateHumanSketchStyleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateHumanSketchStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateHumanSketchStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateHumanSketchStyleResponse) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleResponse) SetHeaders(v map[string]*string) *GenerateHumanSketchStyleResponse {
	s.Headers = v
	return s
}

func (s *GenerateHumanSketchStyleResponse) SetStatusCode(v int32) *GenerateHumanSketchStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateHumanSketchStyleResponse) SetBody(v *GenerateHumanSketchStyleResponseBody) *GenerateHumanSketchStyleResponse {
	s.Body = v
	return s
}

type GetFaceEntityRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s GetFaceEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *GetFaceEntityRequest) SetDbName(v string) *GetFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *GetFaceEntityRequest) SetEntityId(v string) *GetFaceEntityRequest {
	s.EntityId = &v
	return s
}

type GetFaceEntityResponseBody struct {
	Data      *GetFaceEntityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFaceEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponseBody) SetData(v *GetFaceEntityResponseBodyData) *GetFaceEntityResponseBody {
	s.Data = v
	return s
}

func (s *GetFaceEntityResponseBody) SetRequestId(v string) *GetFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

type GetFaceEntityResponseBodyData struct {
	DbName   *string                               `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string                               `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Faces    []*GetFaceEntityResponseBodyDataFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	Labels   *string                               `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s GetFaceEntityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponseBodyData) SetDbName(v string) *GetFaceEntityResponseBodyData {
	s.DbName = &v
	return s
}

func (s *GetFaceEntityResponseBodyData) SetEntityId(v string) *GetFaceEntityResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetFaceEntityResponseBodyData) SetFaces(v []*GetFaceEntityResponseBodyDataFaces) *GetFaceEntityResponseBodyData {
	s.Faces = v
	return s
}

func (s *GetFaceEntityResponseBodyData) SetLabels(v string) *GetFaceEntityResponseBodyData {
	s.Labels = &v
	return s
}

type GetFaceEntityResponseBodyDataFaces struct {
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
}

func (s GetFaceEntityResponseBodyDataFaces) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityResponseBodyDataFaces) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponseBodyDataFaces) SetFaceId(v string) *GetFaceEntityResponseBodyDataFaces {
	s.FaceId = &v
	return s
}

type GetFaceEntityResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFaceEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponse) SetHeaders(v map[string]*string) *GetFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *GetFaceEntityResponse) SetStatusCode(v int32) *GetFaceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFaceEntityResponse) SetBody(v *GetFaceEntityResponseBody) *GetFaceEntityResponse {
	s.Body = v
	return s
}

type GetRealPersonVerificationResultRequest struct {
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s GetRealPersonVerificationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealPersonVerificationResultRequest) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultRequest) SetVerificationToken(v string) *GetRealPersonVerificationResultRequest {
	s.VerificationToken = &v
	return s
}

type GetRealPersonVerificationResultResponseBody struct {
	Data      *GetRealPersonVerificationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealPersonVerificationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealPersonVerificationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultResponseBody) SetData(v *GetRealPersonVerificationResultResponseBodyData) *GetRealPersonVerificationResultResponseBody {
	s.Data = v
	return s
}

func (s *GetRealPersonVerificationResultResponseBody) SetRequestId(v string) *GetRealPersonVerificationResultResponseBody {
	s.RequestId = &v
	return s
}

type GetRealPersonVerificationResultResponseBodyData struct {
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
}

func (s GetRealPersonVerificationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRealPersonVerificationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultResponseBodyData) SetPassed(v bool) *GetRealPersonVerificationResultResponseBodyData {
	s.Passed = &v
	return s
}

type GetRealPersonVerificationResultResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRealPersonVerificationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRealPersonVerificationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealPersonVerificationResultResponse) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultResponse) SetHeaders(v map[string]*string) *GetRealPersonVerificationResultResponse {
	s.Headers = v
	return s
}

func (s *GetRealPersonVerificationResultResponse) SetStatusCode(v int32) *GetRealPersonVerificationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealPersonVerificationResultResponse) SetBody(v *GetRealPersonVerificationResultResponseBody) *GetRealPersonVerificationResultResponse {
	s.Body = v
	return s
}

type HandPostureRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s HandPostureRequest) String() string {
	return tea.Prettify(s)
}

func (s HandPostureRequest) GoString() string {
	return s.String()
}

func (s *HandPostureRequest) SetImageURL(v string) *HandPostureRequest {
	s.ImageURL = &v
	return s
}

type HandPostureAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s HandPostureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s HandPostureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *HandPostureAdvanceRequest) SetImageURLObject(v io.Reader) *HandPostureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type HandPostureResponseBody struct {
	Data      *HandPostureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HandPostureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBody) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBody) SetData(v *HandPostureResponseBodyData) *HandPostureResponseBody {
	s.Data = v
	return s
}

func (s *HandPostureResponseBody) SetRequestId(v string) *HandPostureResponseBody {
	s.RequestId = &v
	return s
}

type HandPostureResponseBodyData struct {
	MetaObject *HandPostureResponseBodyDataMetaObject `json:"MetaObject,omitempty" xml:"MetaObject,omitempty" type:"Struct"`
	Outputs    []*HandPostureResponseBodyDataOutputs  `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyData) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyData) SetMetaObject(v *HandPostureResponseBodyDataMetaObject) *HandPostureResponseBodyData {
	s.MetaObject = v
	return s
}

func (s *HandPostureResponseBodyData) SetOutputs(v []*HandPostureResponseBodyDataOutputs) *HandPostureResponseBodyData {
	s.Outputs = v
	return s
}

type HandPostureResponseBodyDataMetaObject struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s HandPostureResponseBodyDataMetaObject) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataMetaObject) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataMetaObject) SetHeight(v int32) *HandPostureResponseBodyDataMetaObject {
	s.Height = &v
	return s
}

func (s *HandPostureResponseBodyDataMetaObject) SetWidth(v int32) *HandPostureResponseBodyDataMetaObject {
	s.Width = &v
	return s
}

type HandPostureResponseBodyDataOutputs struct {
	HandCount *int32                                       `json:"HandCount,omitempty" xml:"HandCount,omitempty"`
	Results   []*HandPostureResponseBodyDataOutputsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyDataOutputs) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputs) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputs) SetHandCount(v int32) *HandPostureResponseBodyDataOutputs {
	s.HandCount = &v
	return s
}

func (s *HandPostureResponseBodyDataOutputs) SetResults(v []*HandPostureResponseBodyDataOutputsResults) *HandPostureResponseBodyDataOutputs {
	s.Results = v
	return s
}

type HandPostureResponseBodyDataOutputsResults struct {
	Box   *HandPostureResponseBodyDataOutputsResultsBox   `json:"Box,omitempty" xml:"Box,omitempty" type:"Struct"`
	Hands *HandPostureResponseBodyDataOutputsResultsHands `json:"Hands,omitempty" xml:"Hands,omitempty" type:"Struct"`
}

func (s HandPostureResponseBodyDataOutputsResults) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResults) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResults) SetBox(v *HandPostureResponseBodyDataOutputsResultsBox) *HandPostureResponseBodyDataOutputsResults {
	s.Box = v
	return s
}

func (s *HandPostureResponseBodyDataOutputsResults) SetHands(v *HandPostureResponseBodyDataOutputsResultsHands) *HandPostureResponseBodyDataOutputsResults {
	s.Hands = v
	return s
}

type HandPostureResponseBodyDataOutputsResultsBox struct {
	Confident *float32                                                 `json:"Confident,omitempty" xml:"Confident,omitempty"`
	Positions []*HandPostureResponseBodyDataOutputsResultsBoxPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyDataOutputsResultsBox) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsBox) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsBox) SetConfident(v float32) *HandPostureResponseBodyDataOutputsResultsBox {
	s.Confident = &v
	return s
}

func (s *HandPostureResponseBodyDataOutputsResultsBox) SetPositions(v []*HandPostureResponseBodyDataOutputsResultsBoxPositions) *HandPostureResponseBodyDataOutputsResultsBox {
	s.Positions = v
	return s
}

type HandPostureResponseBodyDataOutputsResultsBoxPositions struct {
	Points []*float32 `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyDataOutputsResultsBoxPositions) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsBoxPositions) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsBoxPositions) SetPoints(v []*float32) *HandPostureResponseBodyDataOutputsResultsBoxPositions {
	s.Points = v
	return s
}

type HandPostureResponseBodyDataOutputsResultsHands struct {
	Confident *float32                                                   `json:"Confident,omitempty" xml:"Confident,omitempty"`
	KeyPoints []*HandPostureResponseBodyDataOutputsResultsHandsKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyDataOutputsResultsHands) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsHands) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsHands) SetConfident(v float32) *HandPostureResponseBodyDataOutputsResultsHands {
	s.Confident = &v
	return s
}

func (s *HandPostureResponseBodyDataOutputsResultsHands) SetKeyPoints(v []*HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) *HandPostureResponseBodyDataOutputsResultsHands {
	s.KeyPoints = v
	return s
}

type HandPostureResponseBodyDataOutputsResultsHandsKeyPoints struct {
	Label     *string                                                             `json:"Label,omitempty" xml:"Label,omitempty"`
	Positions []*HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) SetLabel(v string) *HandPostureResponseBodyDataOutputsResultsHandsKeyPoints {
	s.Label = &v
	return s
}

func (s *HandPostureResponseBodyDataOutputsResultsHandsKeyPoints) SetPositions(v []*HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions) *HandPostureResponseBodyDataOutputsResultsHandsKeyPoints {
	s.Positions = v
	return s
}

type HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions struct {
	Points []*float32 `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions) GoString() string {
	return s.String()
}

func (s *HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions) SetPoints(v []*float32) *HandPostureResponseBodyDataOutputsResultsHandsKeyPointsPositions {
	s.Points = v
	return s
}

type HandPostureResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HandPostureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HandPostureResponse) String() string {
	return tea.Prettify(s)
}

func (s HandPostureResponse) GoString() string {
	return s.String()
}

func (s *HandPostureResponse) SetHeaders(v map[string]*string) *HandPostureResponse {
	s.Headers = v
	return s
}

func (s *HandPostureResponse) SetStatusCode(v int32) *HandPostureResponse {
	s.StatusCode = &v
	return s
}

func (s *HandPostureResponse) SetBody(v *HandPostureResponseBody) *HandPostureResponse {
	s.Body = v
	return s
}

type LiquifyFaceRequest struct {
	ImageURL   *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	SlimDegree *float32 `json:"SlimDegree,omitempty" xml:"SlimDegree,omitempty"`
}

func (s LiquifyFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s LiquifyFaceRequest) GoString() string {
	return s.String()
}

func (s *LiquifyFaceRequest) SetImageURL(v string) *LiquifyFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *LiquifyFaceRequest) SetSlimDegree(v float32) *LiquifyFaceRequest {
	s.SlimDegree = &v
	return s
}

type LiquifyFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	SlimDegree     *float32  `json:"SlimDegree,omitempty" xml:"SlimDegree,omitempty"`
}

func (s LiquifyFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s LiquifyFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *LiquifyFaceAdvanceRequest) SetImageURLObject(v io.Reader) *LiquifyFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *LiquifyFaceAdvanceRequest) SetSlimDegree(v float32) *LiquifyFaceAdvanceRequest {
	s.SlimDegree = &v
	return s
}

type LiquifyFaceResponseBody struct {
	Data      *LiquifyFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LiquifyFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LiquifyFaceResponseBody) GoString() string {
	return s.String()
}

func (s *LiquifyFaceResponseBody) SetData(v *LiquifyFaceResponseBodyData) *LiquifyFaceResponseBody {
	s.Data = v
	return s
}

func (s *LiquifyFaceResponseBody) SetRequestId(v string) *LiquifyFaceResponseBody {
	s.RequestId = &v
	return s
}

type LiquifyFaceResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s LiquifyFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LiquifyFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *LiquifyFaceResponseBodyData) SetImageURL(v string) *LiquifyFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

type LiquifyFaceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LiquifyFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LiquifyFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s LiquifyFaceResponse) GoString() string {
	return s.String()
}

func (s *LiquifyFaceResponse) SetHeaders(v map[string]*string) *LiquifyFaceResponse {
	s.Headers = v
	return s
}

func (s *LiquifyFaceResponse) SetStatusCode(v int32) *LiquifyFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *LiquifyFaceResponse) SetBody(v *LiquifyFaceResponseBody) *LiquifyFaceResponse {
	s.Body = v
	return s
}

type ListFaceDbsRequest struct {
	Limit  *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
}

func (s ListFaceDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFaceDbsRequest) GoString() string {
	return s.String()
}

func (s *ListFaceDbsRequest) SetLimit(v int64) *ListFaceDbsRequest {
	s.Limit = &v
	return s
}

func (s *ListFaceDbsRequest) SetOffset(v int64) *ListFaceDbsRequest {
	s.Offset = &v
	return s
}

type ListFaceDbsResponseBody struct {
	Data      *ListFaceDbsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFaceDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceDbsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponseBody) SetData(v *ListFaceDbsResponseBodyData) *ListFaceDbsResponseBody {
	s.Data = v
	return s
}

func (s *ListFaceDbsResponseBody) SetRequestId(v string) *ListFaceDbsResponseBody {
	s.RequestId = &v
	return s
}

type ListFaceDbsResponseBodyData struct {
	DbList []*ListFaceDbsResponseBodyDataDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
}

func (s ListFaceDbsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFaceDbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponseBodyData) SetDbList(v []*ListFaceDbsResponseBodyDataDbList) *ListFaceDbsResponseBodyData {
	s.DbList = v
	return s
}

type ListFaceDbsResponseBodyDataDbList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListFaceDbsResponseBodyDataDbList) String() string {
	return tea.Prettify(s)
}

func (s ListFaceDbsResponseBodyDataDbList) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponseBodyDataDbList) SetName(v string) *ListFaceDbsResponseBodyDataDbList {
	s.Name = &v
	return s
}

type ListFaceDbsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFaceDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceDbsResponse) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponse) SetHeaders(v map[string]*string) *ListFaceDbsResponse {
	s.Headers = v
	return s
}

func (s *ListFaceDbsResponse) SetStatusCode(v int32) *ListFaceDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFaceDbsResponse) SetBody(v *ListFaceDbsResponseBody) *ListFaceDbsResponse {
	s.Body = v
	return s
}

type ListFaceEntitiesRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityIdPrefix *string `json:"EntityIdPrefix,omitempty" xml:"EntityIdPrefix,omitempty"`
	Labels         *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Limit          *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset         *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ListFaceEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesRequest) SetDbName(v string) *ListFaceEntitiesRequest {
	s.DbName = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetEntityIdPrefix(v string) *ListFaceEntitiesRequest {
	s.EntityIdPrefix = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetLabels(v string) *ListFaceEntitiesRequest {
	s.Labels = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetLimit(v int32) *ListFaceEntitiesRequest {
	s.Limit = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetOffset(v int32) *ListFaceEntitiesRequest {
	s.Offset = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetOrder(v string) *ListFaceEntitiesRequest {
	s.Order = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetToken(v string) *ListFaceEntitiesRequest {
	s.Token = &v
	return s
}

type ListFaceEntitiesResponseBody struct {
	Data      *ListFaceEntitiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFaceEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponseBody) SetData(v *ListFaceEntitiesResponseBodyData) *ListFaceEntitiesResponseBody {
	s.Data = v
	return s
}

func (s *ListFaceEntitiesResponseBody) SetRequestId(v string) *ListFaceEntitiesResponseBody {
	s.RequestId = &v
	return s
}

type ListFaceEntitiesResponseBodyData struct {
	Entities   []*ListFaceEntitiesResponseBodyDataEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	Token      *string                                     `json:"Token,omitempty" xml:"Token,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFaceEntitiesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponseBodyData) SetEntities(v []*ListFaceEntitiesResponseBodyDataEntities) *ListFaceEntitiesResponseBodyData {
	s.Entities = v
	return s
}

func (s *ListFaceEntitiesResponseBodyData) SetToken(v string) *ListFaceEntitiesResponseBodyData {
	s.Token = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyData) SetTotalCount(v int32) *ListFaceEntitiesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListFaceEntitiesResponseBodyDataEntities struct {
	CreatedAt *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DbName    *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	FaceCount *int32  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	Labels    *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListFaceEntitiesResponseBodyDataEntities) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesResponseBodyDataEntities) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetCreatedAt(v int64) *ListFaceEntitiesResponseBodyDataEntities {
	s.CreatedAt = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetDbName(v string) *ListFaceEntitiesResponseBodyDataEntities {
	s.DbName = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetEntityId(v string) *ListFaceEntitiesResponseBodyDataEntities {
	s.EntityId = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetFaceCount(v int32) *ListFaceEntitiesResponseBodyDataEntities {
	s.FaceCount = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetLabels(v string) *ListFaceEntitiesResponseBodyDataEntities {
	s.Labels = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetUpdatedAt(v int64) *ListFaceEntitiesResponseBodyDataEntities {
	s.UpdatedAt = &v
	return s
}

type ListFaceEntitiesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFaceEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFaceEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFaceEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponse) SetHeaders(v map[string]*string) *ListFaceEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListFaceEntitiesResponse) SetStatusCode(v int32) *ListFaceEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFaceEntitiesResponse) SetBody(v *ListFaceEntitiesResponseBody) *ListFaceEntitiesResponse {
	s.Body = v
	return s
}

type MergeImageFaceRequest struct {
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s MergeImageFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceRequest) GoString() string {
	return s.String()
}

func (s *MergeImageFaceRequest) SetImageURL(v string) *MergeImageFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *MergeImageFaceRequest) SetTemplateId(v string) *MergeImageFaceRequest {
	s.TemplateId = &v
	return s
}

type MergeImageFaceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	TemplateId     *string   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s MergeImageFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MergeImageFaceAdvanceRequest) SetImageURLObject(v io.Reader) *MergeImageFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *MergeImageFaceAdvanceRequest) SetTemplateId(v string) *MergeImageFaceAdvanceRequest {
	s.TemplateId = &v
	return s
}

type MergeImageFaceResponseBody struct {
	Data      *MergeImageFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MergeImageFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceResponseBody) GoString() string {
	return s.String()
}

func (s *MergeImageFaceResponseBody) SetData(v *MergeImageFaceResponseBodyData) *MergeImageFaceResponseBody {
	s.Data = v
	return s
}

func (s *MergeImageFaceResponseBody) SetRequestId(v string) *MergeImageFaceResponseBody {
	s.RequestId = &v
	return s
}

type MergeImageFaceResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s MergeImageFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MergeImageFaceResponseBodyData) SetImageURL(v string) *MergeImageFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

type MergeImageFaceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MergeImageFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MergeImageFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeImageFaceResponse) GoString() string {
	return s.String()
}

func (s *MergeImageFaceResponse) SetHeaders(v map[string]*string) *MergeImageFaceResponse {
	s.Headers = v
	return s
}

func (s *MergeImageFaceResponse) SetStatusCode(v int32) *MergeImageFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeImageFaceResponse) SetBody(v *MergeImageFaceResponseBody) *MergeImageFaceResponse {
	s.Body = v
	return s
}

type MonitorExaminationRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Type     *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MonitorExaminationRequest) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationRequest) GoString() string {
	return s.String()
}

func (s *MonitorExaminationRequest) SetImageURL(v string) *MonitorExaminationRequest {
	s.ImageURL = &v
	return s
}

func (s *MonitorExaminationRequest) SetType(v int64) *MonitorExaminationRequest {
	s.Type = &v
	return s
}

type MonitorExaminationAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Type           *int64    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MonitorExaminationAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MonitorExaminationAdvanceRequest) SetImageURLObject(v io.Reader) *MonitorExaminationAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *MonitorExaminationAdvanceRequest) SetType(v int64) *MonitorExaminationAdvanceRequest {
	s.Type = &v
	return s
}

type MonitorExaminationResponseBody struct {
	Data      *MonitorExaminationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MonitorExaminationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationResponseBody) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBody) SetData(v *MonitorExaminationResponseBodyData) *MonitorExaminationResponseBody {
	s.Data = v
	return s
}

func (s *MonitorExaminationResponseBody) SetRequestId(v string) *MonitorExaminationResponseBody {
	s.RequestId = &v
	return s
}

type MonitorExaminationResponseBodyData struct {
	ChatScore  *float32                                      `json:"ChatScore,omitempty" xml:"ChatScore,omitempty"`
	FaceInfo   *MonitorExaminationResponseBodyDataFaceInfo   `json:"FaceInfo,omitempty" xml:"FaceInfo,omitempty" type:"Struct"`
	PersonInfo *MonitorExaminationResponseBodyDataPersonInfo `json:"PersonInfo,omitempty" xml:"PersonInfo,omitempty" type:"Struct"`
	Threshold  *float32                                      `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s MonitorExaminationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationResponseBodyData) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyData) SetChatScore(v float32) *MonitorExaminationResponseBodyData {
	s.ChatScore = &v
	return s
}

func (s *MonitorExaminationResponseBodyData) SetFaceInfo(v *MonitorExaminationResponseBodyDataFaceInfo) *MonitorExaminationResponseBodyData {
	s.FaceInfo = v
	return s
}

func (s *MonitorExaminationResponseBodyData) SetPersonInfo(v *MonitorExaminationResponseBodyDataPersonInfo) *MonitorExaminationResponseBodyData {
	s.PersonInfo = v
	return s
}

func (s *MonitorExaminationResponseBodyData) SetThreshold(v float32) *MonitorExaminationResponseBodyData {
	s.Threshold = &v
	return s
}

type MonitorExaminationResponseBodyDataFaceInfo struct {
	Completeness *float32                                        `json:"Completeness,omitempty" xml:"Completeness,omitempty"`
	FaceNumber   *int64                                          `json:"FaceNumber,omitempty" xml:"FaceNumber,omitempty"`
	Pose         *MonitorExaminationResponseBodyDataFaceInfoPose `json:"Pose,omitempty" xml:"Pose,omitempty" type:"Struct"`
}

func (s MonitorExaminationResponseBodyDataFaceInfo) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataFaceInfo) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) SetCompleteness(v float32) *MonitorExaminationResponseBodyDataFaceInfo {
	s.Completeness = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) SetFaceNumber(v int64) *MonitorExaminationResponseBodyDataFaceInfo {
	s.FaceNumber = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfo) SetPose(v *MonitorExaminationResponseBodyDataFaceInfoPose) *MonitorExaminationResponseBodyDataFaceInfo {
	s.Pose = v
	return s
}

type MonitorExaminationResponseBodyDataFaceInfoPose struct {
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s MonitorExaminationResponseBodyDataFaceInfoPose) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataFaceInfoPose) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) SetPitch(v float32) *MonitorExaminationResponseBodyDataFaceInfoPose {
	s.Pitch = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) SetRoll(v float32) *MonitorExaminationResponseBodyDataFaceInfoPose {
	s.Roll = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataFaceInfoPose) SetYaw(v float32) *MonitorExaminationResponseBodyDataFaceInfoPose {
	s.Yaw = &v
	return s
}

type MonitorExaminationResponseBodyDataPersonInfo struct {
	CellPhone    *MonitorExaminationResponseBodyDataPersonInfoCellPhone `json:"CellPhone,omitempty" xml:"CellPhone,omitempty" type:"Struct"`
	EarPhone     *MonitorExaminationResponseBodyDataPersonInfoEarPhone  `json:"EarPhone,omitempty" xml:"EarPhone,omitempty" type:"Struct"`
	PersonNumber *int64                                                 `json:"PersonNumber,omitempty" xml:"PersonNumber,omitempty"`
}

func (s MonitorExaminationResponseBodyDataPersonInfo) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataPersonInfo) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) SetCellPhone(v *MonitorExaminationResponseBodyDataPersonInfoCellPhone) *MonitorExaminationResponseBodyDataPersonInfo {
	s.CellPhone = v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) SetEarPhone(v *MonitorExaminationResponseBodyDataPersonInfoEarPhone) *MonitorExaminationResponseBodyDataPersonInfo {
	s.EarPhone = v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfo) SetPersonNumber(v int64) *MonitorExaminationResponseBodyDataPersonInfo {
	s.PersonNumber = &v
	return s
}

type MonitorExaminationResponseBodyDataPersonInfoCellPhone struct {
	Score     *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s MonitorExaminationResponseBodyDataPersonInfoCellPhone) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataPersonInfoCellPhone) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataPersonInfoCellPhone) SetScore(v float32) *MonitorExaminationResponseBodyDataPersonInfoCellPhone {
	s.Score = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfoCellPhone) SetThreshold(v float32) *MonitorExaminationResponseBodyDataPersonInfoCellPhone {
	s.Threshold = &v
	return s
}

type MonitorExaminationResponseBodyDataPersonInfoEarPhone struct {
	Score     *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s MonitorExaminationResponseBodyDataPersonInfoEarPhone) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationResponseBodyDataPersonInfoEarPhone) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponseBodyDataPersonInfoEarPhone) SetScore(v float32) *MonitorExaminationResponseBodyDataPersonInfoEarPhone {
	s.Score = &v
	return s
}

func (s *MonitorExaminationResponseBodyDataPersonInfoEarPhone) SetThreshold(v float32) *MonitorExaminationResponseBodyDataPersonInfoEarPhone {
	s.Threshold = &v
	return s
}

type MonitorExaminationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MonitorExaminationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MonitorExaminationResponse) String() string {
	return tea.Prettify(s)
}

func (s MonitorExaminationResponse) GoString() string {
	return s.String()
}

func (s *MonitorExaminationResponse) SetHeaders(v map[string]*string) *MonitorExaminationResponse {
	s.Headers = v
	return s
}

func (s *MonitorExaminationResponse) SetStatusCode(v int32) *MonitorExaminationResponse {
	s.StatusCode = &v
	return s
}

func (s *MonitorExaminationResponse) SetBody(v *MonitorExaminationResponseBody) *MonitorExaminationResponse {
	s.Body = v
	return s
}

type PedestrianDetectAttributeRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s PedestrianDetectAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeRequest) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeRequest) SetImageURL(v string) *PedestrianDetectAttributeRequest {
	s.ImageURL = &v
	return s
}

type PedestrianDetectAttributeAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s PedestrianDetectAttributeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeAdvanceRequest) SetImageURLObject(v io.Reader) *PedestrianDetectAttributeAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type PedestrianDetectAttributeResponseBody struct {
	Data      *PedestrianDetectAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PedestrianDetectAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBody) SetData(v *PedestrianDetectAttributeResponseBodyData) *PedestrianDetectAttributeResponseBody {
	s.Data = v
	return s
}

func (s *PedestrianDetectAttributeResponseBody) SetRequestId(v string) *PedestrianDetectAttributeResponseBody {
	s.RequestId = &v
	return s
}

type PedestrianDetectAttributeResponseBodyData struct {
	Attributes   []*PedestrianDetectAttributeResponseBodyDataAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Boxes        []*PedestrianDetectAttributeResponseBodyDataBoxes      `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Height       *int64                                                 `json:"Height,omitempty" xml:"Height,omitempty"`
	PersonNumber *int32                                                 `json:"PersonNumber,omitempty" xml:"PersonNumber,omitempty"`
	Width        *int64                                                 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyData) SetAttributes(v []*PedestrianDetectAttributeResponseBodyDataAttributes) *PedestrianDetectAttributeResponseBodyData {
	s.Attributes = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetBoxes(v []*PedestrianDetectAttributeResponseBodyDataBoxes) *PedestrianDetectAttributeResponseBodyData {
	s.Boxes = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetHeight(v int64) *PedestrianDetectAttributeResponseBodyData {
	s.Height = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetPersonNumber(v int32) *PedestrianDetectAttributeResponseBodyData {
	s.PersonNumber = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetWidth(v int64) *PedestrianDetectAttributeResponseBodyData {
	s.Width = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributes struct {
	Age         *PedestrianDetectAttributeResponseBodyDataAttributesAge         `json:"Age,omitempty" xml:"Age,omitempty" type:"Struct"`
	Backpack    *PedestrianDetectAttributeResponseBodyDataAttributesBackpack    `json:"Backpack,omitempty" xml:"Backpack,omitempty" type:"Struct"`
	Gender      *PedestrianDetectAttributeResponseBodyDataAttributesGender      `json:"Gender,omitempty" xml:"Gender,omitempty" type:"Struct"`
	Glasses     *PedestrianDetectAttributeResponseBodyDataAttributesGlasses     `json:"Glasses,omitempty" xml:"Glasses,omitempty" type:"Struct"`
	Handbag     *PedestrianDetectAttributeResponseBodyDataAttributesHandbag     `json:"Handbag,omitempty" xml:"Handbag,omitempty" type:"Struct"`
	Hat         *PedestrianDetectAttributeResponseBodyDataAttributesHat         `json:"Hat,omitempty" xml:"Hat,omitempty" type:"Struct"`
	LowerColor  *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor  `json:"LowerColor,omitempty" xml:"LowerColor,omitempty" type:"Struct"`
	LowerWear   *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear   `json:"LowerWear,omitempty" xml:"LowerWear,omitempty" type:"Struct"`
	Orient      *PedestrianDetectAttributeResponseBodyDataAttributesOrient      `json:"Orient,omitempty" xml:"Orient,omitempty" type:"Struct"`
	ShoulderBag *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag `json:"ShoulderBag,omitempty" xml:"ShoulderBag,omitempty" type:"Struct"`
	UpperColor  *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor  `json:"UpperColor,omitempty" xml:"UpperColor,omitempty" type:"Struct"`
	UpperWear   *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear   `json:"UpperWear,omitempty" xml:"UpperWear,omitempty" type:"Struct"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributes) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributes) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetAge(v *PedestrianDetectAttributeResponseBodyDataAttributesAge) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Age = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetBackpack(v *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Backpack = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetGender(v *PedestrianDetectAttributeResponseBodyDataAttributesGender) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Gender = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetGlasses(v *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Glasses = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetHandbag(v *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Handbag = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetHat(v *PedestrianDetectAttributeResponseBodyDataAttributesHat) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Hat = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetLowerColor(v *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.LowerColor = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetLowerWear(v *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.LowerWear = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetOrient(v *PedestrianDetectAttributeResponseBodyDataAttributesOrient) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Orient = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetShoulderBag(v *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.ShoulderBag = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetUpperColor(v *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.UpperColor = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetUpperWear(v *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.UpperWear = v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesAge struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesAge) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesAge) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesAge) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesAge {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesAge) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesAge {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesBackpack struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesBackpack) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesBackpack) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesBackpack {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesBackpack {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesGender struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGender) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGender) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGender) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesGender {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGender) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesGender {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesGlasses struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGlasses) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGlasses) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesGlasses {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesGlasses {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesHandbag struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHandbag) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHandbag) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesHandbag {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesHandbag {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesHat struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHat) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHat) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHat) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesHat {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHat) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesHat {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesLowerColor struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesLowerWear struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesOrient struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesOrient) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesOrient) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesOrient) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesOrient {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesOrient) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesOrient {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesUpperColor struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataAttributesUpperWear struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear {
	s.Score = &v
	return s
}

type PedestrianDetectAttributeResponseBodyDataBoxes struct {
	BottomRightX *float32 `json:"BottomRightX,omitempty" xml:"BottomRightX,omitempty"`
	BottomRightY *float32 `json:"BottomRightY,omitempty" xml:"BottomRightY,omitempty"`
	Score        *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	TopLeftX     *float32 `json:"TopLeftX,omitempty" xml:"TopLeftX,omitempty"`
	TopLeftY     *float32 `json:"TopLeftY,omitempty" xml:"TopLeftY,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataBoxes) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataBoxes) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetBottomRightX(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.BottomRightX = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetBottomRightY(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.BottomRightY = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetTopLeftX(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.TopLeftX = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetTopLeftY(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.TopLeftY = &v
	return s
}

type PedestrianDetectAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PedestrianDetectAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PedestrianDetectAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s PedestrianDetectAttributeResponse) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponse) SetHeaders(v map[string]*string) *PedestrianDetectAttributeResponse {
	s.Headers = v
	return s
}

func (s *PedestrianDetectAttributeResponse) SetStatusCode(v int32) *PedestrianDetectAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *PedestrianDetectAttributeResponse) SetBody(v *PedestrianDetectAttributeResponseBody) *PedestrianDetectAttributeResponse {
	s.Body = v
	return s
}

type QueryFaceImageTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryFaceImageTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateRequest) SetTemplateId(v string) *QueryFaceImageTemplateRequest {
	s.TemplateId = &v
	return s
}

type QueryFaceImageTemplateResponseBody struct {
	Data      *QueryFaceImageTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFaceImageTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBody) SetData(v *QueryFaceImageTemplateResponseBodyData) *QueryFaceImageTemplateResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceImageTemplateResponseBody) SetRequestId(v string) *QueryFaceImageTemplateResponseBody {
	s.RequestId = &v
	return s
}

type QueryFaceImageTemplateResponseBodyData struct {
	Elements []*QueryFaceImageTemplateResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s QueryFaceImageTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBodyData) SetElements(v []*QueryFaceImageTemplateResponseBodyDataElements) *QueryFaceImageTemplateResponseBodyData {
	s.Elements = v
	return s
}

type QueryFaceImageTemplateResponseBodyDataElements struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceImageTemplateResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetCreateTime(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.CreateTime = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetTemplateId(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.TemplateId = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetTemplateURL(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.TemplateURL = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetUpdateTime(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.UpdateTime = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetUserId(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.UserId = &v
	return s
}

type QueryFaceImageTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryFaceImageTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFaceImageTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFaceImageTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponse) SetHeaders(v map[string]*string) *QueryFaceImageTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceImageTemplateResponse) SetStatusCode(v int32) *QueryFaceImageTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceImageTemplateResponse) SetBody(v *QueryFaceImageTemplateResponseBody) *QueryFaceImageTemplateResponse {
	s.Body = v
	return s
}

type RecognizeActionRequest struct {
	Type      *int32                           `json:"Type,omitempty" xml:"Type,omitempty"`
	URLList   []*RecognizeActionRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
	VideoData *string                          `json:"VideoData,omitempty" xml:"VideoData,omitempty"`
	VideoUrl  *string                          `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeActionRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionRequest) GoString() string {
	return s.String()
}

func (s *RecognizeActionRequest) SetType(v int32) *RecognizeActionRequest {
	s.Type = &v
	return s
}

func (s *RecognizeActionRequest) SetURLList(v []*RecognizeActionRequestURLList) *RecognizeActionRequest {
	s.URLList = v
	return s
}

func (s *RecognizeActionRequest) SetVideoData(v string) *RecognizeActionRequest {
	s.VideoData = &v
	return s
}

func (s *RecognizeActionRequest) SetVideoUrl(v string) *RecognizeActionRequest {
	s.VideoUrl = &v
	return s
}

type RecognizeActionRequestURLList struct {
	URL       *string `json:"URL,omitempty" xml:"URL,omitempty"`
	ImageData *string `json:"imageData,omitempty" xml:"imageData,omitempty"`
}

func (s RecognizeActionRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionRequestURLList) GoString() string {
	return s.String()
}

func (s *RecognizeActionRequestURLList) SetURL(v string) *RecognizeActionRequestURLList {
	s.URL = &v
	return s
}

func (s *RecognizeActionRequestURLList) SetImageData(v string) *RecognizeActionRequestURLList {
	s.ImageData = &v
	return s
}

type RecognizeActionAdvanceRequest struct {
	Type           *int32                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	URLList        []*RecognizeActionAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
	VideoData      *string                                 `json:"VideoData,omitempty" xml:"VideoData,omitempty"`
	VideoUrlObject io.Reader                               `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeActionAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeActionAdvanceRequest) SetType(v int32) *RecognizeActionAdvanceRequest {
	s.Type = &v
	return s
}

func (s *RecognizeActionAdvanceRequest) SetURLList(v []*RecognizeActionAdvanceRequestURLList) *RecognizeActionAdvanceRequest {
	s.URLList = v
	return s
}

func (s *RecognizeActionAdvanceRequest) SetVideoData(v string) *RecognizeActionAdvanceRequest {
	s.VideoData = &v
	return s
}

func (s *RecognizeActionAdvanceRequest) SetVideoUrlObject(v io.Reader) *RecognizeActionAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type RecognizeActionAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
	ImageData *string   `json:"imageData,omitempty" xml:"imageData,omitempty"`
}

func (s RecognizeActionAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *RecognizeActionAdvanceRequestURLList) SetURLObject(v io.Reader) *RecognizeActionAdvanceRequestURLList {
	s.URLObject = v
	return s
}

func (s *RecognizeActionAdvanceRequestURLList) SetImageData(v string) *RecognizeActionAdvanceRequestURLList {
	s.ImageData = &v
	return s
}

type RecognizeActionResponseBody struct {
	Data      *RecognizeActionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBody) SetData(v *RecognizeActionResponseBodyData) *RecognizeActionResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeActionResponseBody) SetRequestId(v string) *RecognizeActionResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeActionResponseBodyData struct {
	Elements []*RecognizeActionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeActionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBodyData) SetElements(v []*RecognizeActionResponseBodyDataElements) *RecognizeActionResponseBodyData {
	s.Elements = v
	return s
}

type RecognizeActionResponseBodyDataElements struct {
	Boxes     []*RecognizeActionResponseBodyDataElementsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Labels    []*string                                       `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Scores    []*float32                                      `json:"Scores,omitempty" xml:"Scores,omitempty" type:"Repeated"`
	Timestamp *int32                                          `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s RecognizeActionResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBodyDataElements) SetBoxes(v []*RecognizeActionResponseBodyDataElementsBoxes) *RecognizeActionResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) SetLabels(v []*string) *RecognizeActionResponseBodyDataElements {
	s.Labels = v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) SetScores(v []*float32) *RecognizeActionResponseBodyDataElements {
	s.Scores = v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) SetTimestamp(v int32) *RecognizeActionResponseBodyDataElements {
	s.Timestamp = &v
	return s
}

type RecognizeActionResponseBodyDataElementsBoxes struct {
	Box []*int32 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
}

func (s RecognizeActionResponseBodyDataElementsBoxes) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBodyDataElementsBoxes) SetBox(v []*int32) *RecognizeActionResponseBodyDataElementsBoxes {
	s.Box = v
	return s
}

type RecognizeActionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeActionResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeActionResponse) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponse) SetHeaders(v map[string]*string) *RecognizeActionResponse {
	s.Headers = v
	return s
}

func (s *RecognizeActionResponse) SetStatusCode(v int32) *RecognizeActionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeActionResponse) SetBody(v *RecognizeActionResponseBody) *RecognizeActionResponse {
	s.Body = v
	return s
}

type RecognizeExpressionRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeExpressionRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionRequest) SetImageURL(v string) *RecognizeExpressionRequest {
	s.ImageURL = &v
	return s
}

type RecognizeExpressionAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeExpressionAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeExpressionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeExpressionResponseBody struct {
	Data      *RecognizeExpressionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeExpressionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBody) SetData(v *RecognizeExpressionResponseBodyData) *RecognizeExpressionResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeExpressionResponseBody) SetRequestId(v string) *RecognizeExpressionResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeExpressionResponseBodyData struct {
	Elements []*RecognizeExpressionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeExpressionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBodyData) SetElements(v []*RecognizeExpressionResponseBodyDataElements) *RecognizeExpressionResponseBodyData {
	s.Elements = v
	return s
}

type RecognizeExpressionResponseBodyDataElements struct {
	Expression      *string                                                   `json:"Expression,omitempty" xml:"Expression,omitempty"`
	FaceProbability *float32                                                  `json:"FaceProbability,omitempty" xml:"FaceProbability,omitempty"`
	FaceRectangle   *RecognizeExpressionResponseBodyDataElementsFaceRectangle `json:"FaceRectangle,omitempty" xml:"FaceRectangle,omitempty" type:"Struct"`
}

func (s RecognizeExpressionResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBodyDataElements) SetExpression(v string) *RecognizeExpressionResponseBodyDataElements {
	s.Expression = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElements) SetFaceProbability(v float32) *RecognizeExpressionResponseBodyDataElements {
	s.FaceProbability = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElements) SetFaceRectangle(v *RecognizeExpressionResponseBodyDataElementsFaceRectangle) *RecognizeExpressionResponseBodyDataElements {
	s.FaceRectangle = v
	return s
}

type RecognizeExpressionResponseBodyDataElementsFaceRectangle struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeExpressionResponseBodyDataElementsFaceRectangle) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponseBodyDataElementsFaceRectangle) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetHeight(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Height = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetLeft(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Left = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetTop(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Top = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetWidth(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Width = &v
	return s
}

type RecognizeExpressionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeExpressionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeExpressionResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExpressionResponse) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponse) SetHeaders(v map[string]*string) *RecognizeExpressionResponse {
	s.Headers = v
	return s
}

func (s *RecognizeExpressionResponse) SetStatusCode(v int32) *RecognizeExpressionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeExpressionResponse) SetBody(v *RecognizeExpressionResponseBody) *RecognizeExpressionResponse {
	s.Body = v
	return s
}

type RecognizeFaceRequest struct {
	Age           *bool   `json:"Age,omitempty" xml:"Age,omitempty"`
	Beauty        *bool   `json:"Beauty,omitempty" xml:"Beauty,omitempty"`
	Expression    *bool   `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Gender        *bool   `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Glass         *bool   `json:"Glass,omitempty" xml:"Glass,omitempty"`
	Hat           *bool   `json:"Hat,omitempty" xml:"Hat,omitempty"`
	ImageURL      *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Mask          *bool   `json:"Mask,omitempty" xml:"Mask,omitempty"`
	MaxFaceNumber *int64  `json:"MaxFaceNumber,omitempty" xml:"MaxFaceNumber,omitempty"`
	Quality       *bool   `json:"Quality,omitempty" xml:"Quality,omitempty"`
}

func (s RecognizeFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFaceRequest) SetAge(v bool) *RecognizeFaceRequest {
	s.Age = &v
	return s
}

func (s *RecognizeFaceRequest) SetBeauty(v bool) *RecognizeFaceRequest {
	s.Beauty = &v
	return s
}

func (s *RecognizeFaceRequest) SetExpression(v bool) *RecognizeFaceRequest {
	s.Expression = &v
	return s
}

func (s *RecognizeFaceRequest) SetGender(v bool) *RecognizeFaceRequest {
	s.Gender = &v
	return s
}

func (s *RecognizeFaceRequest) SetGlass(v bool) *RecognizeFaceRequest {
	s.Glass = &v
	return s
}

func (s *RecognizeFaceRequest) SetHat(v bool) *RecognizeFaceRequest {
	s.Hat = &v
	return s
}

func (s *RecognizeFaceRequest) SetImageURL(v string) *RecognizeFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeFaceRequest) SetMask(v bool) *RecognizeFaceRequest {
	s.Mask = &v
	return s
}

func (s *RecognizeFaceRequest) SetMaxFaceNumber(v int64) *RecognizeFaceRequest {
	s.MaxFaceNumber = &v
	return s
}

func (s *RecognizeFaceRequest) SetQuality(v bool) *RecognizeFaceRequest {
	s.Quality = &v
	return s
}

type RecognizeFaceAdvanceRequest struct {
	Age            *bool     `json:"Age,omitempty" xml:"Age,omitempty"`
	Beauty         *bool     `json:"Beauty,omitempty" xml:"Beauty,omitempty"`
	Expression     *bool     `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Gender         *bool     `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Glass          *bool     `json:"Glass,omitempty" xml:"Glass,omitempty"`
	Hat            *bool     `json:"Hat,omitempty" xml:"Hat,omitempty"`
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Mask           *bool     `json:"Mask,omitempty" xml:"Mask,omitempty"`
	MaxFaceNumber  *int64    `json:"MaxFaceNumber,omitempty" xml:"MaxFaceNumber,omitempty"`
	Quality        *bool     `json:"Quality,omitempty" xml:"Quality,omitempty"`
}

func (s RecognizeFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFaceAdvanceRequest) SetAge(v bool) *RecognizeFaceAdvanceRequest {
	s.Age = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetBeauty(v bool) *RecognizeFaceAdvanceRequest {
	s.Beauty = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetExpression(v bool) *RecognizeFaceAdvanceRequest {
	s.Expression = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetGender(v bool) *RecognizeFaceAdvanceRequest {
	s.Gender = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetGlass(v bool) *RecognizeFaceAdvanceRequest {
	s.Glass = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetHat(v bool) *RecognizeFaceAdvanceRequest {
	s.Hat = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetMask(v bool) *RecognizeFaceAdvanceRequest {
	s.Mask = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetMaxFaceNumber(v int64) *RecognizeFaceAdvanceRequest {
	s.MaxFaceNumber = &v
	return s
}

func (s *RecognizeFaceAdvanceRequest) SetQuality(v bool) *RecognizeFaceAdvanceRequest {
	s.Quality = &v
	return s
}

type RecognizeFaceResponseBody struct {
	Data      *RecognizeFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBody) SetData(v *RecognizeFaceResponseBodyData) *RecognizeFaceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeFaceResponseBody) SetRequestId(v string) *RecognizeFaceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeFaceResponseBodyData struct {
	// 1
	AgeList []*int32 `json:"AgeList,omitempty" xml:"AgeList,omitempty" type:"Repeated"`
	// 1
	BeautyList         []*float32 `json:"BeautyList,omitempty" xml:"BeautyList,omitempty" type:"Repeated"`
	DenseFeatureLength *int32     `json:"DenseFeatureLength,omitempty" xml:"DenseFeatureLength,omitempty"`
	// 1
	DenseFeatures []*string `json:"DenseFeatures,omitempty" xml:"DenseFeatures,omitempty" type:"Repeated"`
	// 1
	Expressions []*int32 `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	FaceCount   *int32   `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	// 1
	FaceProbabilityList []*float32 `json:"FaceProbabilityList,omitempty" xml:"FaceProbabilityList,omitempty" type:"Repeated"`
	// 1
	FaceRectangles []*int32 `json:"FaceRectangles,omitempty" xml:"FaceRectangles,omitempty" type:"Repeated"`
	// 1
	GenderList []*int32 `json:"GenderList,omitempty" xml:"GenderList,omitempty" type:"Repeated"`
	// 1
	Glasses []*int32 `json:"Glasses,omitempty" xml:"Glasses,omitempty" type:"Repeated"`
	// 1
	HatList       []*int32 `json:"HatList,omitempty" xml:"HatList,omitempty" type:"Repeated"`
	LandmarkCount *int32   `json:"LandmarkCount,omitempty" xml:"LandmarkCount,omitempty"`
	// 1
	Landmarks []*float32 `json:"Landmarks,omitempty" xml:"Landmarks,omitempty" type:"Repeated"`
	// 1
	Masks []*int64 `json:"Masks,omitempty" xml:"Masks,omitempty" type:"Repeated"`
	// 1
	PoseList []*float32 `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
	// 1
	Pupils    []*float32                              `json:"Pupils,omitempty" xml:"Pupils,omitempty" type:"Repeated"`
	Qualities *RecognizeFaceResponseBodyDataQualities `json:"Qualities,omitempty" xml:"Qualities,omitempty" type:"Struct"`
}

func (s RecognizeFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBodyData) SetAgeList(v []*int32) *RecognizeFaceResponseBodyData {
	s.AgeList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetBeautyList(v []*float32) *RecognizeFaceResponseBodyData {
	s.BeautyList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetDenseFeatureLength(v int32) *RecognizeFaceResponseBodyData {
	s.DenseFeatureLength = &v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetDenseFeatures(v []*string) *RecognizeFaceResponseBodyData {
	s.DenseFeatures = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetExpressions(v []*int32) *RecognizeFaceResponseBodyData {
	s.Expressions = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetFaceCount(v int32) *RecognizeFaceResponseBodyData {
	s.FaceCount = &v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetFaceProbabilityList(v []*float32) *RecognizeFaceResponseBodyData {
	s.FaceProbabilityList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetFaceRectangles(v []*int32) *RecognizeFaceResponseBodyData {
	s.FaceRectangles = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetGenderList(v []*int32) *RecognizeFaceResponseBodyData {
	s.GenderList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetGlasses(v []*int32) *RecognizeFaceResponseBodyData {
	s.Glasses = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetHatList(v []*int32) *RecognizeFaceResponseBodyData {
	s.HatList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetLandmarkCount(v int32) *RecognizeFaceResponseBodyData {
	s.LandmarkCount = &v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetLandmarks(v []*float32) *RecognizeFaceResponseBodyData {
	s.Landmarks = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetMasks(v []*int64) *RecognizeFaceResponseBodyData {
	s.Masks = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetPoseList(v []*float32) *RecognizeFaceResponseBodyData {
	s.PoseList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetPupils(v []*float32) *RecognizeFaceResponseBodyData {
	s.Pupils = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetQualities(v *RecognizeFaceResponseBodyDataQualities) *RecognizeFaceResponseBodyData {
	s.Qualities = v
	return s
}

type RecognizeFaceResponseBodyDataQualities struct {
	// 1
	BlurList []*float32 `json:"BlurList,omitempty" xml:"BlurList,omitempty" type:"Repeated"`
	// 1
	FnfList []*float32 `json:"FnfList,omitempty" xml:"FnfList,omitempty" type:"Repeated"`
	// 1
	GlassList []*float32 `json:"GlassList,omitempty" xml:"GlassList,omitempty" type:"Repeated"`
	// 1
	IlluList []*float32 `json:"IlluList,omitempty" xml:"IlluList,omitempty" type:"Repeated"`
	// 1
	MaskList []*float32 `json:"MaskList,omitempty" xml:"MaskList,omitempty" type:"Repeated"`
	// 1
	NoiseList []*float32 `json:"NoiseList,omitempty" xml:"NoiseList,omitempty" type:"Repeated"`
	// 1
	PoseList []*float32 `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
	// 1
	ScoreList []*float32 `json:"ScoreList,omitempty" xml:"ScoreList,omitempty" type:"Repeated"`
}

func (s RecognizeFaceResponseBodyDataQualities) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponseBodyDataQualities) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBodyDataQualities) SetBlurList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.BlurList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetFnfList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.FnfList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetGlassList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.GlassList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetIlluList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.IlluList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetMaskList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.MaskList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetNoiseList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.NoiseList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetPoseList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.PoseList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetScoreList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.ScoreList = v
	return s
}

type RecognizeFaceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFaceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponse) SetHeaders(v map[string]*string) *RecognizeFaceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFaceResponse) SetStatusCode(v int32) *RecognizeFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeFaceResponse) SetBody(v *RecognizeFaceResponseBody) *RecognizeFaceResponse {
	s.Body = v
	return s
}

type RecognizeHandGestureRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GestureType *string `json:"GestureType,omitempty" xml:"GestureType,omitempty"`
	ImageURL    *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeHandGestureRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandGestureRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHandGestureRequest) SetAppId(v string) *RecognizeHandGestureRequest {
	s.AppId = &v
	return s
}

func (s *RecognizeHandGestureRequest) SetGestureType(v string) *RecognizeHandGestureRequest {
	s.GestureType = &v
	return s
}

func (s *RecognizeHandGestureRequest) SetImageURL(v string) *RecognizeHandGestureRequest {
	s.ImageURL = &v
	return s
}

type RecognizeHandGestureAdvanceRequest struct {
	AppId          *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GestureType    *string   `json:"GestureType,omitempty" xml:"GestureType,omitempty"`
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeHandGestureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandGestureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHandGestureAdvanceRequest) SetAppId(v string) *RecognizeHandGestureAdvanceRequest {
	s.AppId = &v
	return s
}

func (s *RecognizeHandGestureAdvanceRequest) SetGestureType(v string) *RecognizeHandGestureAdvanceRequest {
	s.GestureType = &v
	return s
}

func (s *RecognizeHandGestureAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeHandGestureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeHandGestureResponseBody struct {
	Data      *RecognizeHandGestureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHandGestureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandGestureResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHandGestureResponseBody) SetData(v *RecognizeHandGestureResponseBodyData) *RecognizeHandGestureResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeHandGestureResponseBody) SetRequestId(v string) *RecognizeHandGestureResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHandGestureResponseBodyData struct {
	Height *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	Score  *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Width  *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int64   `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int64   `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeHandGestureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandGestureResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeHandGestureResponseBodyData) SetHeight(v int64) *RecognizeHandGestureResponseBodyData {
	s.Height = &v
	return s
}

func (s *RecognizeHandGestureResponseBodyData) SetScore(v float32) *RecognizeHandGestureResponseBodyData {
	s.Score = &v
	return s
}

func (s *RecognizeHandGestureResponseBodyData) SetType(v string) *RecognizeHandGestureResponseBodyData {
	s.Type = &v
	return s
}

func (s *RecognizeHandGestureResponseBodyData) SetWidth(v int64) *RecognizeHandGestureResponseBodyData {
	s.Width = &v
	return s
}

func (s *RecognizeHandGestureResponseBodyData) SetX(v int64) *RecognizeHandGestureResponseBodyData {
	s.X = &v
	return s
}

func (s *RecognizeHandGestureResponseBodyData) SetY(v int64) *RecognizeHandGestureResponseBodyData {
	s.Y = &v
	return s
}

type RecognizeHandGestureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeHandGestureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeHandGestureResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandGestureResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHandGestureResponse) SetHeaders(v map[string]*string) *RecognizeHandGestureResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHandGestureResponse) SetStatusCode(v int32) *RecognizeHandGestureResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHandGestureResponse) SetBody(v *RecognizeHandGestureResponseBody) *RecognizeHandGestureResponse {
	s.Body = v
	return s
}

type RecognizePublicFaceRequest struct {
	// 1
	Task []*RecognizePublicFaceRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s RecognizePublicFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceRequest) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceRequest) SetTask(v []*RecognizePublicFaceRequestTask) *RecognizePublicFaceRequest {
	s.Task = v
	return s
}

type RecognizePublicFaceRequestTask struct {
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizePublicFaceRequestTask) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceRequestTask) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceRequestTask) SetImageData(v string) *RecognizePublicFaceRequestTask {
	s.ImageData = &v
	return s
}

func (s *RecognizePublicFaceRequestTask) SetImageURL(v string) *RecognizePublicFaceRequestTask {
	s.ImageURL = &v
	return s
}

type RecognizePublicFaceAdvanceRequest struct {
	// 1
	Task []*RecognizePublicFaceAdvanceRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s RecognizePublicFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceAdvanceRequest) SetTask(v []*RecognizePublicFaceAdvanceRequestTask) *RecognizePublicFaceAdvanceRequest {
	s.Task = v
	return s
}

type RecognizePublicFaceAdvanceRequestTask struct {
	ImageData      *string   `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizePublicFaceAdvanceRequestTask) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceAdvanceRequestTask) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceAdvanceRequestTask) SetImageData(v string) *RecognizePublicFaceAdvanceRequestTask {
	s.ImageData = &v
	return s
}

func (s *RecognizePublicFaceAdvanceRequestTask) SetImageURLObject(v io.Reader) *RecognizePublicFaceAdvanceRequestTask {
	s.ImageURLObject = v
	return s
}

type RecognizePublicFaceResponseBody struct {
	Data      *RecognizePublicFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePublicFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBody) SetData(v *RecognizePublicFaceResponseBodyData) *RecognizePublicFaceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizePublicFaceResponseBody) SetRequestId(v string) *RecognizePublicFaceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizePublicFaceResponseBodyData struct {
	Elements []*RecognizePublicFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizePublicFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyData) SetElements(v []*RecognizePublicFaceResponseBodyDataElements) *RecognizePublicFaceResponseBodyData {
	s.Elements = v
	return s
}

type RecognizePublicFaceResponseBodyDataElements struct {
	ImageURL *string                                               `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Results  []*RecognizePublicFaceResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	TaskId   *string                                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElements) SetImageURL(v string) *RecognizePublicFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElements) SetResults(v []*RecognizePublicFaceResponseBodyDataElementsResults) *RecognizePublicFaceResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElements) SetTaskId(v string) *RecognizePublicFaceResponseBodyDataElements {
	s.TaskId = &v
	return s
}

type RecognizePublicFaceResponseBodyDataElementsResults struct {
	Label      *string                                                         `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *float32                                                        `json:"Rate,omitempty" xml:"Rate,omitempty"`
	SubResults []*RecognizePublicFaceResponseBodyDataElementsResultsSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Repeated"`
	Suggestion *string                                                         `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElementsResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetLabel(v string) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetRate(v float32) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetSubResults(v []*RecognizePublicFaceResponseBodyDataElementsResultsSubResults) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.SubResults = v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetSuggestion(v string) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

type RecognizePublicFaceResponseBodyDataElementsResultsSubResults struct {
	Faces []*RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	H     *float32                                                             `json:"H,omitempty" xml:"H,omitempty"`
	W     *float32                                                             `json:"W,omitempty" xml:"W,omitempty"`
	X     *float32                                                             `json:"X,omitempty" xml:"X,omitempty"`
	Y     *float32                                                             `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResults) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetFaces(v []*RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.Faces = v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetH(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.H = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetW(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.W = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetX(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.X = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetY(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.Y = &v
	return s
}

type RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces struct {
	Id   *string  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) SetId(v string) *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	s.Id = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) SetName(v string) *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	s.Name = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) SetRate(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	s.Rate = &v
	return s
}

type RecognizePublicFaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizePublicFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizePublicFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePublicFaceResponse) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponse) SetHeaders(v map[string]*string) *RecognizePublicFaceResponse {
	s.Headers = v
	return s
}

func (s *RecognizePublicFaceResponse) SetStatusCode(v int32) *RecognizePublicFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePublicFaceResponse) SetBody(v *RecognizePublicFaceResponseBody) *RecognizePublicFaceResponse {
	s.Body = v
	return s
}

type RetouchBodyRequest struct {
	ImageURL       *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	LengthenDegree *float32 `json:"LengthenDegree,omitempty" xml:"LengthenDegree,omitempty"`
	SlimDegree     *float32 `json:"SlimDegree,omitempty" xml:"SlimDegree,omitempty"`
}

func (s RetouchBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s RetouchBodyRequest) GoString() string {
	return s.String()
}

func (s *RetouchBodyRequest) SetImageURL(v string) *RetouchBodyRequest {
	s.ImageURL = &v
	return s
}

func (s *RetouchBodyRequest) SetLengthenDegree(v float32) *RetouchBodyRequest {
	s.LengthenDegree = &v
	return s
}

func (s *RetouchBodyRequest) SetSlimDegree(v float32) *RetouchBodyRequest {
	s.SlimDegree = &v
	return s
}

type RetouchBodyAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	LengthenDegree *float32  `json:"LengthenDegree,omitempty" xml:"LengthenDegree,omitempty"`
	SlimDegree     *float32  `json:"SlimDegree,omitempty" xml:"SlimDegree,omitempty"`
}

func (s RetouchBodyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RetouchBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RetouchBodyAdvanceRequest) SetImageURLObject(v io.Reader) *RetouchBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RetouchBodyAdvanceRequest) SetLengthenDegree(v float32) *RetouchBodyAdvanceRequest {
	s.LengthenDegree = &v
	return s
}

func (s *RetouchBodyAdvanceRequest) SetSlimDegree(v float32) *RetouchBodyAdvanceRequest {
	s.SlimDegree = &v
	return s
}

type RetouchBodyResponseBody struct {
	Data      *RetouchBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetouchBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetouchBodyResponseBody) GoString() string {
	return s.String()
}

func (s *RetouchBodyResponseBody) SetData(v *RetouchBodyResponseBodyData) *RetouchBodyResponseBody {
	s.Data = v
	return s
}

func (s *RetouchBodyResponseBody) SetRequestId(v string) *RetouchBodyResponseBody {
	s.RequestId = &v
	return s
}

type RetouchBodyResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RetouchBodyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RetouchBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetouchBodyResponseBodyData) SetImageURL(v string) *RetouchBodyResponseBodyData {
	s.ImageURL = &v
	return s
}

type RetouchBodyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RetouchBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetouchBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s RetouchBodyResponse) GoString() string {
	return s.String()
}

func (s *RetouchBodyResponse) SetHeaders(v map[string]*string) *RetouchBodyResponse {
	s.Headers = v
	return s
}

func (s *RetouchBodyResponse) SetStatusCode(v int32) *RetouchBodyResponse {
	s.StatusCode = &v
	return s
}

func (s *RetouchBodyResponse) SetBody(v *RetouchBodyResponseBody) *RetouchBodyResponse {
	s.Body = v
	return s
}

type RetouchSkinRequest struct {
	ImageURL        *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RetouchDegree   *float32 `json:"RetouchDegree,omitempty" xml:"RetouchDegree,omitempty"`
	WhiteningDegree *float32 `json:"WhiteningDegree,omitempty" xml:"WhiteningDegree,omitempty"`
}

func (s RetouchSkinRequest) String() string {
	return tea.Prettify(s)
}

func (s RetouchSkinRequest) GoString() string {
	return s.String()
}

func (s *RetouchSkinRequest) SetImageURL(v string) *RetouchSkinRequest {
	s.ImageURL = &v
	return s
}

func (s *RetouchSkinRequest) SetRetouchDegree(v float32) *RetouchSkinRequest {
	s.RetouchDegree = &v
	return s
}

func (s *RetouchSkinRequest) SetWhiteningDegree(v float32) *RetouchSkinRequest {
	s.WhiteningDegree = &v
	return s
}

type RetouchSkinAdvanceRequest struct {
	ImageURLObject  io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RetouchDegree   *float32  `json:"RetouchDegree,omitempty" xml:"RetouchDegree,omitempty"`
	WhiteningDegree *float32  `json:"WhiteningDegree,omitempty" xml:"WhiteningDegree,omitempty"`
}

func (s RetouchSkinAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RetouchSkinAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RetouchSkinAdvanceRequest) SetImageURLObject(v io.Reader) *RetouchSkinAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RetouchSkinAdvanceRequest) SetRetouchDegree(v float32) *RetouchSkinAdvanceRequest {
	s.RetouchDegree = &v
	return s
}

func (s *RetouchSkinAdvanceRequest) SetWhiteningDegree(v float32) *RetouchSkinAdvanceRequest {
	s.WhiteningDegree = &v
	return s
}

type RetouchSkinResponseBody struct {
	Data      *RetouchSkinResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetouchSkinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetouchSkinResponseBody) GoString() string {
	return s.String()
}

func (s *RetouchSkinResponseBody) SetData(v *RetouchSkinResponseBodyData) *RetouchSkinResponseBody {
	s.Data = v
	return s
}

func (s *RetouchSkinResponseBody) SetRequestId(v string) *RetouchSkinResponseBody {
	s.RequestId = &v
	return s
}

type RetouchSkinResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RetouchSkinResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RetouchSkinResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetouchSkinResponseBodyData) SetImageURL(v string) *RetouchSkinResponseBodyData {
	s.ImageURL = &v
	return s
}

type RetouchSkinResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RetouchSkinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetouchSkinResponse) String() string {
	return tea.Prettify(s)
}

func (s RetouchSkinResponse) GoString() string {
	return s.String()
}

func (s *RetouchSkinResponse) SetHeaders(v map[string]*string) *RetouchSkinResponse {
	s.Headers = v
	return s
}

func (s *RetouchSkinResponse) SetStatusCode(v int32) *RetouchSkinResponse {
	s.StatusCode = &v
	return s
}

func (s *RetouchSkinResponse) SetBody(v *RetouchSkinResponseBody) *RetouchSkinResponse {
	s.Body = v
	return s
}

type SearchFaceRequest struct {
	DbName                *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbNames               *string  `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	ImageUrl              *string  `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Limit                 *int32   `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MaxFaceNum            *int64   `json:"MaxFaceNum,omitempty" xml:"MaxFaceNum,omitempty"`
	QualityScoreThreshold *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
}

func (s SearchFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceRequest) GoString() string {
	return s.String()
}

func (s *SearchFaceRequest) SetDbName(v string) *SearchFaceRequest {
	s.DbName = &v
	return s
}

func (s *SearchFaceRequest) SetDbNames(v string) *SearchFaceRequest {
	s.DbNames = &v
	return s
}

func (s *SearchFaceRequest) SetImageUrl(v string) *SearchFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *SearchFaceRequest) SetLimit(v int32) *SearchFaceRequest {
	s.Limit = &v
	return s
}

func (s *SearchFaceRequest) SetMaxFaceNum(v int64) *SearchFaceRequest {
	s.MaxFaceNum = &v
	return s
}

func (s *SearchFaceRequest) SetQualityScoreThreshold(v float32) *SearchFaceRequest {
	s.QualityScoreThreshold = &v
	return s
}

type SearchFaceAdvanceRequest struct {
	DbName                *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbNames               *string   `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	ImageUrlObject        io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Limit                 *int32    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MaxFaceNum            *int64    `json:"MaxFaceNum,omitempty" xml:"MaxFaceNum,omitempty"`
	QualityScoreThreshold *float32  `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
}

func (s SearchFaceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SearchFaceAdvanceRequest) SetDbName(v string) *SearchFaceAdvanceRequest {
	s.DbName = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetDbNames(v string) *SearchFaceAdvanceRequest {
	s.DbNames = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetImageUrlObject(v io.Reader) *SearchFaceAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *SearchFaceAdvanceRequest) SetLimit(v int32) *SearchFaceAdvanceRequest {
	s.Limit = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetMaxFaceNum(v int64) *SearchFaceAdvanceRequest {
	s.MaxFaceNum = &v
	return s
}

func (s *SearchFaceAdvanceRequest) SetQualityScoreThreshold(v float32) *SearchFaceAdvanceRequest {
	s.QualityScoreThreshold = &v
	return s
}

type SearchFaceResponseBody struct {
	Data      *SearchFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBody) SetData(v *SearchFaceResponseBodyData) *SearchFaceResponseBody {
	s.Data = v
	return s
}

func (s *SearchFaceResponseBody) SetRequestId(v string) *SearchFaceResponseBody {
	s.RequestId = &v
	return s
}

type SearchFaceResponseBodyData struct {
	MatchList []*SearchFaceResponseBodyDataMatchList `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
}

func (s SearchFaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyData) SetMatchList(v []*SearchFaceResponseBodyDataMatchList) *SearchFaceResponseBodyData {
	s.MatchList = v
	return s
}

type SearchFaceResponseBodyDataMatchList struct {
	FaceItems     []*SearchFaceResponseBodyDataMatchListFaceItems `json:"FaceItems,omitempty" xml:"FaceItems,omitempty" type:"Repeated"`
	Location      *SearchFaceResponseBodyDataMatchListLocation    `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	QualitieScore *float32                                        `json:"QualitieScore,omitempty" xml:"QualitieScore,omitempty"`
}

func (s SearchFaceResponseBodyDataMatchList) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyDataMatchList) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataMatchList) SetFaceItems(v []*SearchFaceResponseBodyDataMatchListFaceItems) *SearchFaceResponseBodyDataMatchList {
	s.FaceItems = v
	return s
}

func (s *SearchFaceResponseBodyDataMatchList) SetLocation(v *SearchFaceResponseBodyDataMatchListLocation) *SearchFaceResponseBodyDataMatchList {
	s.Location = v
	return s
}

func (s *SearchFaceResponseBodyDataMatchList) SetQualitieScore(v float32) *SearchFaceResponseBodyDataMatchList {
	s.QualitieScore = &v
	return s
}

type SearchFaceResponseBodyDataMatchListFaceItems struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	DbName     *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId   *string  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ExtraData  *string  `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	FaceId     *string  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	Score      *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchFaceResponseBodyDataMatchListFaceItems) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyDataMatchListFaceItems) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetConfidence(v float32) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.Confidence = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetDbName(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.DbName = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetEntityId(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.EntityId = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetExtraData(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.ExtraData = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetFaceId(v string) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.FaceId = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListFaceItems) SetScore(v float32) *SearchFaceResponseBodyDataMatchListFaceItems {
	s.Score = &v
	return s
}

type SearchFaceResponseBodyDataMatchListLocation struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SearchFaceResponseBodyDataMatchListLocation) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponseBodyDataMatchListLocation) GoString() string {
	return s.String()
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetHeight(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.Height = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetWidth(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.Width = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetX(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.X = &v
	return s
}

func (s *SearchFaceResponseBodyDataMatchListLocation) SetY(v int32) *SearchFaceResponseBodyDataMatchListLocation {
	s.Y = &v
	return s
}

type SearchFaceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFaceResponse) GoString() string {
	return s.String()
}

func (s *SearchFaceResponse) SetHeaders(v map[string]*string) *SearchFaceResponse {
	s.Headers = v
	return s
}

func (s *SearchFaceResponse) SetStatusCode(v int32) *SearchFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFaceResponse) SetBody(v *SearchFaceResponseBody) *SearchFaceResponse {
	s.Body = v
	return s
}

type SwapFacialFeaturesRequest struct {
	EditPart        *string `json:"EditPart,omitempty" xml:"EditPart,omitempty"`
	SourceImageData []byte  `json:"SourceImageData,omitempty" xml:"SourceImageData,omitempty"`
	SourceImageURL  *string `json:"SourceImageURL,omitempty" xml:"SourceImageURL,omitempty"`
	TargetImageData []byte  `json:"TargetImageData,omitempty" xml:"TargetImageData,omitempty"`
	TargetImageURL  *string `json:"TargetImageURL,omitempty" xml:"TargetImageURL,omitempty"`
}

func (s SwapFacialFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesRequest) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesRequest) SetEditPart(v string) *SwapFacialFeaturesRequest {
	s.EditPart = &v
	return s
}

func (s *SwapFacialFeaturesRequest) SetSourceImageData(v []byte) *SwapFacialFeaturesRequest {
	s.SourceImageData = v
	return s
}

func (s *SwapFacialFeaturesRequest) SetSourceImageURL(v string) *SwapFacialFeaturesRequest {
	s.SourceImageURL = &v
	return s
}

func (s *SwapFacialFeaturesRequest) SetTargetImageData(v []byte) *SwapFacialFeaturesRequest {
	s.TargetImageData = v
	return s
}

func (s *SwapFacialFeaturesRequest) SetTargetImageURL(v string) *SwapFacialFeaturesRequest {
	s.TargetImageURL = &v
	return s
}

type SwapFacialFeaturesAdvanceRequest struct {
	EditPart             *string   `json:"EditPart,omitempty" xml:"EditPart,omitempty"`
	SourceImageData      []byte    `json:"SourceImageData,omitempty" xml:"SourceImageData,omitempty"`
	SourceImageURLObject io.Reader `json:"SourceImageURL,omitempty" xml:"SourceImageURL,omitempty"`
	TargetImageData      []byte    `json:"TargetImageData,omitempty" xml:"TargetImageData,omitempty"`
	TargetImageURLObject io.Reader `json:"TargetImageURL,omitempty" xml:"TargetImageURL,omitempty"`
}

func (s SwapFacialFeaturesAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesAdvanceRequest) SetEditPart(v string) *SwapFacialFeaturesAdvanceRequest {
	s.EditPart = &v
	return s
}

func (s *SwapFacialFeaturesAdvanceRequest) SetSourceImageData(v []byte) *SwapFacialFeaturesAdvanceRequest {
	s.SourceImageData = v
	return s
}

func (s *SwapFacialFeaturesAdvanceRequest) SetSourceImageURLObject(v io.Reader) *SwapFacialFeaturesAdvanceRequest {
	s.SourceImageURLObject = v
	return s
}

func (s *SwapFacialFeaturesAdvanceRequest) SetTargetImageData(v []byte) *SwapFacialFeaturesAdvanceRequest {
	s.TargetImageData = v
	return s
}

func (s *SwapFacialFeaturesAdvanceRequest) SetTargetImageURLObject(v io.Reader) *SwapFacialFeaturesAdvanceRequest {
	s.TargetImageURLObject = v
	return s
}

type SwapFacialFeaturesResponseBody struct {
	Data      *SwapFacialFeaturesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwapFacialFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesResponseBody) SetData(v *SwapFacialFeaturesResponseBodyData) *SwapFacialFeaturesResponseBody {
	s.Data = v
	return s
}

func (s *SwapFacialFeaturesResponseBody) SetRequestId(v string) *SwapFacialFeaturesResponseBody {
	s.RequestId = &v
	return s
}

type SwapFacialFeaturesResponseBodyData struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SwapFacialFeaturesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesResponseBodyData) SetImageURL(v string) *SwapFacialFeaturesResponseBodyData {
	s.ImageURL = &v
	return s
}

type SwapFacialFeaturesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwapFacialFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwapFacialFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s SwapFacialFeaturesResponse) GoString() string {
	return s.String()
}

func (s *SwapFacialFeaturesResponse) SetHeaders(v map[string]*string) *SwapFacialFeaturesResponse {
	s.Headers = v
	return s
}

func (s *SwapFacialFeaturesResponse) SetStatusCode(v int32) *SwapFacialFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *SwapFacialFeaturesResponse) SetBody(v *SwapFacialFeaturesResponseBody) *SwapFacialFeaturesResponse {
	s.Body = v
	return s
}

type UpdateFaceEntityRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Labels   *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateFaceEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceEntityRequest) SetDbName(v string) *UpdateFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *UpdateFaceEntityRequest) SetEntityId(v string) *UpdateFaceEntityRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateFaceEntityRequest) SetLabels(v string) *UpdateFaceEntityRequest {
	s.Labels = &v
	return s
}

type UpdateFaceEntityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFaceEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceEntityResponseBody) SetRequestId(v string) *UpdateFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFaceEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFaceEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceEntityResponse) SetHeaders(v map[string]*string) *UpdateFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceEntityResponse) SetStatusCode(v int32) *UpdateFaceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFaceEntityResponse) SetBody(v *UpdateFaceEntityResponseBody) *UpdateFaceEntityResponse {
	s.Body = v
	return s
}

type VerifyFaceMaskRequest struct {
	ImageData []byte  `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RefData   []byte  `json:"RefData,omitempty" xml:"RefData,omitempty"`
	RefUrl    *string `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
}

func (s VerifyFaceMaskRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskRequest) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskRequest) SetImageData(v []byte) *VerifyFaceMaskRequest {
	s.ImageData = v
	return s
}

func (s *VerifyFaceMaskRequest) SetImageURL(v string) *VerifyFaceMaskRequest {
	s.ImageURL = &v
	return s
}

func (s *VerifyFaceMaskRequest) SetRefData(v []byte) *VerifyFaceMaskRequest {
	s.RefData = v
	return s
}

func (s *VerifyFaceMaskRequest) SetRefUrl(v string) *VerifyFaceMaskRequest {
	s.RefUrl = &v
	return s
}

type VerifyFaceMaskAdvanceRequest struct {
	ImageData      []byte    `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RefData        []byte    `json:"RefData,omitempty" xml:"RefData,omitempty"`
	RefUrlObject   io.Reader `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
}

func (s VerifyFaceMaskAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskAdvanceRequest) SetImageData(v []byte) *VerifyFaceMaskAdvanceRequest {
	s.ImageData = v
	return s
}

func (s *VerifyFaceMaskAdvanceRequest) SetImageURLObject(v io.Reader) *VerifyFaceMaskAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *VerifyFaceMaskAdvanceRequest) SetRefData(v []byte) *VerifyFaceMaskAdvanceRequest {
	s.RefData = v
	return s
}

func (s *VerifyFaceMaskAdvanceRequest) SetRefUrlObject(v io.Reader) *VerifyFaceMaskAdvanceRequest {
	s.RefUrlObject = v
	return s
}

type VerifyFaceMaskResponseBody struct {
	Data      *VerifyFaceMaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyFaceMaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskResponseBody) SetData(v *VerifyFaceMaskResponseBodyData) *VerifyFaceMaskResponseBody {
	s.Data = v
	return s
}

func (s *VerifyFaceMaskResponseBody) SetRequestId(v string) *VerifyFaceMaskResponseBody {
	s.RequestId = &v
	return s
}

type VerifyFaceMaskResponseBodyData struct {
	Confidence   *float32   `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Mask         *int32     `json:"Mask,omitempty" xml:"Mask,omitempty"`
	MaskRef      *int32     `json:"MaskRef,omitempty" xml:"MaskRef,omitempty"`
	Rectangle    []*int32   `json:"Rectangle,omitempty" xml:"Rectangle,omitempty" type:"Repeated"`
	RectangleRef []*int32   `json:"RectangleRef,omitempty" xml:"RectangleRef,omitempty" type:"Repeated"`
	Thresholds   []*float32 `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Repeated"`
}

func (s VerifyFaceMaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskResponseBodyData) SetConfidence(v float32) *VerifyFaceMaskResponseBodyData {
	s.Confidence = &v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetMask(v int32) *VerifyFaceMaskResponseBodyData {
	s.Mask = &v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetMaskRef(v int32) *VerifyFaceMaskResponseBodyData {
	s.MaskRef = &v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetRectangle(v []*int32) *VerifyFaceMaskResponseBodyData {
	s.Rectangle = v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetRectangleRef(v []*int32) *VerifyFaceMaskResponseBodyData {
	s.RectangleRef = v
	return s
}

func (s *VerifyFaceMaskResponseBodyData) SetThresholds(v []*float32) *VerifyFaceMaskResponseBodyData {
	s.Thresholds = v
	return s
}

type VerifyFaceMaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyFaceMaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyFaceMaskResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyFaceMaskResponse) GoString() string {
	return s.String()
}

func (s *VerifyFaceMaskResponse) SetHeaders(v map[string]*string) *VerifyFaceMaskResponse {
	s.Headers = v
	return s
}

func (s *VerifyFaceMaskResponse) SetStatusCode(v int32) *VerifyFaceMaskResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyFaceMaskResponse) SetBody(v *VerifyFaceMaskResponseBody) *VerifyFaceMaskResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("facebody"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceWithOptions(request *AddFaceRequest, runtime *util.RuntimeOptions) (_result *AddFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraData)) {
		body["ExtraData"] = request.ExtraData
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.QualityScoreThreshold)) {
		body["QualityScoreThreshold"] = request.QualityScoreThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.SimilarityScoreThresholdBetweenEntity)) {
		body["SimilarityScoreThresholdBetweenEntity"] = request.SimilarityScoreThresholdBetweenEntity
	}

	if !tea.BoolValue(util.IsUnset(request.SimilarityScoreThresholdInEntity)) {
		body["SimilarityScoreThresholdInEntity"] = request.SimilarityScoreThresholdInEntity
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFace(request *AddFaceRequest) (_result *AddFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceResponse{}
	_body, _err := client.AddFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceAdvance(request *AddFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	addFaceReq := &AddFaceRequest{}
	openapiutil.Convert(request, addFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		addFaceReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	addFaceResp, _err := client.AddFaceWithOptions(addFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addFaceResp
	return _result, _err
}

func (client *Client) AddFaceEntityWithOptions(request *AddFaceEntityRequest, runtime *util.RuntimeOptions) (_result *AddFaceEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceEntity"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceEntity(request *AddFaceEntityRequest) (_result *AddFaceEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceEntityResponse{}
	_body, _err := client.AddFaceEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceImageTemplateWithOptions(request *AddFaceImageTemplateRequest, runtime *util.RuntimeOptions) (_result *AddFaceImageTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaceImageTemplate"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFaceImageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFaceImageTemplate(request *AddFaceImageTemplateRequest) (_result *AddFaceImageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFaceImageTemplateResponse{}
	_body, _err := client.AddFaceImageTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFaceImageTemplateAdvance(request *AddFaceImageTemplateAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddFaceImageTemplateResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	addFaceImageTemplateReq := &AddFaceImageTemplateRequest{}
	openapiutil.Convert(request, addFaceImageTemplateReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		addFaceImageTemplateReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	addFaceImageTemplateResp, _err := client.AddFaceImageTemplateWithOptions(addFaceImageTemplateReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addFaceImageTemplateResp
	return _result, _err
}

func (client *Client) BatchAddFacesWithOptions(tmpReq *BatchAddFacesRequest, runtime *util.RuntimeOptions) (_result *BatchAddFacesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchAddFacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Faces)) {
		request.FacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Faces, tea.String("Faces"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.FacesShrink)) {
		body["Faces"] = request.FacesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.QualityScoreThreshold)) {
		body["QualityScoreThreshold"] = request.QualityScoreThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.SimilarityScoreThresholdBetweenEntity)) {
		body["SimilarityScoreThresholdBetweenEntity"] = request.SimilarityScoreThresholdBetweenEntity
	}

	if !tea.BoolValue(util.IsUnset(request.SimilarityScoreThresholdInEntity)) {
		body["SimilarityScoreThresholdInEntity"] = request.SimilarityScoreThresholdInEntity
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchAddFaces"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchAddFaces(request *BatchAddFacesRequest) (_result *BatchAddFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchAddFacesResponse{}
	_body, _err := client.BatchAddFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BeautifyBodyWithOptions(tmpReq *BeautifyBodyRequest, runtime *util.RuntimeOptions) (_result *BeautifyBodyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BeautifyBodyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AgeRange)) {
		request.AgeRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgeRange, tea.String("AgeRange"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.BodyBoxes)) {
		request.BodyBoxesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyBoxes, tea.String("BodyBoxes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.FaceList)) {
		request.FaceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FaceList, tea.String("FaceList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PoseList)) {
		request.PoseListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PoseList, tea.String("PoseList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgeRangeShrink)) {
		body["AgeRange"] = request.AgeRangeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.BodyBoxesShrink)) {
		body["BodyBoxes"] = request.BodyBoxesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Custom)) {
		body["Custom"] = request.Custom
	}

	if !tea.BoolValue(util.IsUnset(request.FaceListShrink)) {
		body["FaceList"] = request.FaceListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FemaleLiquifyDegree)) {
		body["FemaleLiquifyDegree"] = request.FemaleLiquifyDegree
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.IsPregnant)) {
		body["IsPregnant"] = request.IsPregnant
	}

	if !tea.BoolValue(util.IsUnset(request.LengthenDegree)) {
		body["LengthenDegree"] = request.LengthenDegree
	}

	if !tea.BoolValue(util.IsUnset(request.MaleLiquifyDegree)) {
		body["MaleLiquifyDegree"] = request.MaleLiquifyDegree
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalHeight)) {
		body["OriginalHeight"] = request.OriginalHeight
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalWidth)) {
		body["OriginalWidth"] = request.OriginalWidth
	}

	if !tea.BoolValue(util.IsUnset(request.PoseListShrink)) {
		body["PoseList"] = request.PoseListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BeautifyBody"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BeautifyBodyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BeautifyBody(request *BeautifyBodyRequest) (_result *BeautifyBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BeautifyBodyResponse{}
	_body, _err := client.BeautifyBodyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BeautifyBodyAdvance(request *BeautifyBodyAdvanceRequest, runtime *util.RuntimeOptions) (_result *BeautifyBodyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	beautifyBodyReq := &BeautifyBodyRequest{}
	openapiutil.Convert(request, beautifyBodyReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		beautifyBodyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	beautifyBodyResp, _err := client.BeautifyBodyWithOptions(beautifyBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = beautifyBodyResp
	return _result, _err
}

func (client *Client) BlurFaceWithOptions(request *BlurFaceRequest, runtime *util.RuntimeOptions) (_result *BlurFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BlurFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BlurFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BlurFace(request *BlurFaceRequest) (_result *BlurFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BlurFaceResponse{}
	_body, _err := client.BlurFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BlurFaceAdvance(request *BlurFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *BlurFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	blurFaceReq := &BlurFaceRequest{}
	openapiutil.Convert(request, blurFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		blurFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	blurFaceResp, _err := client.BlurFaceWithOptions(blurFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = blurFaceResp
	return _result, _err
}

func (client *Client) BodyPostureWithOptions(request *BodyPostureRequest, runtime *util.RuntimeOptions) (_result *BodyPostureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BodyPosture"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BodyPostureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BodyPosture(request *BodyPostureRequest) (_result *BodyPostureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BodyPostureResponse{}
	_body, _err := client.BodyPostureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BodyPostureAdvance(request *BodyPostureAdvanceRequest, runtime *util.RuntimeOptions) (_result *BodyPostureResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	bodyPostureReq := &BodyPostureRequest{}
	openapiutil.Convert(request, bodyPostureReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		bodyPostureReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	bodyPostureResp, _err := client.BodyPostureWithOptions(bodyPostureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = bodyPostureResp
	return _result, _err
}

func (client *Client) CompareFaceWithOptions(request *CompareFaceRequest, runtime *util.RuntimeOptions) (_result *CompareFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageDataA)) {
		body["ImageDataA"] = request.ImageDataA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageDataB)) {
		body["ImageDataB"] = request.ImageDataB
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURLA)) {
		body["ImageURLA"] = request.ImageURLA
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURLB)) {
		body["ImageURLB"] = request.ImageURLB
	}

	if !tea.BoolValue(util.IsUnset(request.QualityScoreThreshold)) {
		body["QualityScoreThreshold"] = request.QualityScoreThreshold
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CompareFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CompareFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompareFace(request *CompareFaceRequest) (_result *CompareFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareFaceResponse{}
	_body, _err := client.CompareFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareFaceAdvance(request *CompareFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *CompareFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	compareFaceReq := &CompareFaceRequest{}
	openapiutil.Convert(request, compareFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLAObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLAObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		compareFaceReq.ImageURLA = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURLBObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLBObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		compareFaceReq.ImageURLB = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	compareFaceResp, _err := client.CompareFaceWithOptions(compareFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = compareFaceResp
	return _result, _err
}

func (client *Client) CountCrowdWithOptions(request *CountCrowdRequest, runtime *util.RuntimeOptions) (_result *CountCrowdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.IsShow)) {
		body["IsShow"] = request.IsShow
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CountCrowd"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountCrowdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountCrowd(request *CountCrowdRequest) (_result *CountCrowdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountCrowdResponse{}
	_body, _err := client.CountCrowdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountCrowdAdvance(request *CountCrowdAdvanceRequest, runtime *util.RuntimeOptions) (_result *CountCrowdResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	countCrowdReq := &CountCrowdRequest{}
	openapiutil.Convert(request, countCrowdReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		countCrowdReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	countCrowdResp, _err := client.CountCrowdWithOptions(countCrowdReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = countCrowdResp
	return _result, _err
}

func (client *Client) CreateFaceDbWithOptions(request *CreateFaceDbRequest, runtime *util.RuntimeOptions) (_result *CreateFaceDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFaceDb"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFaceDbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFaceDb(request *CreateFaceDbRequest) (_result *CreateFaceDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFaceDbResponse{}
	_body, _err := client.CreateFaceDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceWithOptions(request *DeleteFaceRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		body["FaceId"] = request.FaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFace(request *DeleteFaceRequest) (_result *DeleteFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceResponse{}
	_body, _err := client.DeleteFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceDbWithOptions(request *DeleteFaceDbRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceDb"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceDbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceDb(request *DeleteFaceDbRequest) (_result *DeleteFaceDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceDbResponse{}
	_body, _err := client.DeleteFaceDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceEntityWithOptions(request *DeleteFaceEntityRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceEntity"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceEntity(request *DeleteFaceEntityRequest) (_result *DeleteFaceEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceEntityResponse{}
	_body, _err := client.DeleteFaceEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFaceImageTemplateWithOptions(request *DeleteFaceImageTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteFaceImageTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaceImageTemplate"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFaceImageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaceImageTemplate(request *DeleteFaceImageTemplateRequest) (_result *DeleteFaceImageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFaceImageTemplateResponse{}
	_body, _err := client.DeleteFaceImageTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectBodyCountWithOptions(request *DetectBodyCountRequest, runtime *util.RuntimeOptions) (_result *DetectBodyCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectBodyCount"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectBodyCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectBodyCount(request *DetectBodyCountRequest) (_result *DetectBodyCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectBodyCountResponse{}
	_body, _err := client.DetectBodyCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectBodyCountAdvance(request *DetectBodyCountAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectBodyCountResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectBodyCountReq := &DetectBodyCountRequest{}
	openapiutil.Convert(request, detectBodyCountReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectBodyCountReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectBodyCountResp, _err := client.DetectBodyCountWithOptions(detectBodyCountReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectBodyCountResp
	return _result, _err
}

func (client *Client) DetectCelebrityWithOptions(request *DetectCelebrityRequest, runtime *util.RuntimeOptions) (_result *DetectCelebrityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectCelebrity"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectCelebrityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectCelebrity(request *DetectCelebrityRequest) (_result *DetectCelebrityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectCelebrityResponse{}
	_body, _err := client.DetectCelebrityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectCelebrityAdvance(request *DetectCelebrityAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectCelebrityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectCelebrityReq := &DetectCelebrityRequest{}
	openapiutil.Convert(request, detectCelebrityReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectCelebrityReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectCelebrityResp, _err := client.DetectCelebrityWithOptions(detectCelebrityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectCelebrityResp
	return _result, _err
}

func (client *Client) DetectChefCapWithOptions(request *DetectChefCapRequest, runtime *util.RuntimeOptions) (_result *DetectChefCapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectChefCap"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectChefCapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectChefCap(request *DetectChefCapRequest) (_result *DetectChefCapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectChefCapResponse{}
	_body, _err := client.DetectChefCapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectChefCapAdvance(request *DetectChefCapAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectChefCapResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectChefCapReq := &DetectChefCapRequest{}
	openapiutil.Convert(request, detectChefCapReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectChefCapReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectChefCapResp, _err := client.DetectChefCapWithOptions(detectChefCapReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectChefCapResp
	return _result, _err
}

func (client *Client) DetectFaceWithOptions(request *DetectFaceRequest, runtime *util.RuntimeOptions) (_result *DetectFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Landmark)) {
		body["Landmark"] = request.Landmark
	}

	if !tea.BoolValue(util.IsUnset(request.MaxFaceNumber)) {
		body["MaxFaceNumber"] = request.MaxFaceNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Pose)) {
		body["Pose"] = request.Pose
	}

	if !tea.BoolValue(util.IsUnset(request.Quality)) {
		body["Quality"] = request.Quality
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectFace(request *DetectFaceRequest) (_result *DetectFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectFaceResponse{}
	_body, _err := client.DetectFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectFaceAdvance(request *DetectFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectFaceReq := &DetectFaceRequest{}
	openapiutil.Convert(request, detectFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectFaceResp, _err := client.DetectFaceWithOptions(detectFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectFaceResp
	return _result, _err
}

func (client *Client) DetectIPCPedestrianWithOptions(request *DetectIPCPedestrianRequest, runtime *util.RuntimeOptions) (_result *DetectIPCPedestrianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["Height"] = request.Height
	}

	if !tea.BoolValue(util.IsUnset(request.ImageData)) {
		body["ImageData"] = request.ImageData
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["Width"] = request.Width
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectIPCPedestrian"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectIPCPedestrianResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectIPCPedestrian(request *DetectIPCPedestrianRequest) (_result *DetectIPCPedestrianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectIPCPedestrianResponse{}
	_body, _err := client.DetectIPCPedestrianWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectIPCPedestrianAdvance(request *DetectIPCPedestrianAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectIPCPedestrianResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectIPCPedestrianReq := &DetectIPCPedestrianRequest{}
	openapiutil.Convert(request, detectIPCPedestrianReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectIPCPedestrianReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectIPCPedestrianResp, _err := client.DetectIPCPedestrianWithOptions(detectIPCPedestrianReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectIPCPedestrianResp
	return _result, _err
}

func (client *Client) DetectLivingFaceWithOptions(request *DetectLivingFaceRequest, runtime *util.RuntimeOptions) (_result *DetectLivingFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tasks)) {
		body["Tasks"] = request.Tasks
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectLivingFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectLivingFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectLivingFace(request *DetectLivingFaceRequest) (_result *DetectLivingFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectLivingFaceResponse{}
	_body, _err := client.DetectLivingFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectLivingFaceAdvance(request *DetectLivingFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectLivingFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectLivingFaceReq := &DetectLivingFaceRequest{}
	openapiutil.Convert(request, detectLivingFaceReq)
	if !tea.BoolValue(util.IsUnset(request.Tasks)) {
		i0 := tea.Int(0)
		for _, item0 := range request.Tasks {
			if !tea.BoolValue(util.IsUnset(item0.ImageURLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.ImageURLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := detectLivingFaceReq.Tasks[tea.IntValue(i0)]
				tmp.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	detectLivingFaceResp, _err := client.DetectLivingFaceWithOptions(detectLivingFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectLivingFaceResp
	return _result, _err
}

func (client *Client) DetectPedestrianWithOptions(request *DetectPedestrianRequest, runtime *util.RuntimeOptions) (_result *DetectPedestrianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectPedestrian"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectPedestrianResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectPedestrian(request *DetectPedestrianRequest) (_result *DetectPedestrianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectPedestrianResponse{}
	_body, _err := client.DetectPedestrianWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectPedestrianAdvance(request *DetectPedestrianAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectPedestrianResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectPedestrianReq := &DetectPedestrianRequest{}
	openapiutil.Convert(request, detectPedestrianReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectPedestrianReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectPedestrianResp, _err := client.DetectPedestrianWithOptions(detectPedestrianReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectPedestrianResp
	return _result, _err
}

func (client *Client) DetectPedestrianIntrusionWithOptions(tmpReq *DetectPedestrianIntrusionRequest, runtime *util.RuntimeOptions) (_result *DetectPedestrianIntrusionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectPedestrianIntrusionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DetectRegion)) {
		request.DetectRegionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DetectRegion, tea.String("DetectRegion"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DetectRegionShrink)) {
		body["DetectRegion"] = request.DetectRegionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.RegionType)) {
		body["RegionType"] = request.RegionType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectPedestrianIntrusion"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectPedestrianIntrusionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectPedestrianIntrusion(request *DetectPedestrianIntrusionRequest) (_result *DetectPedestrianIntrusionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectPedestrianIntrusionResponse{}
	_body, _err := client.DetectPedestrianIntrusionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectPedestrianIntrusionAdvance(request *DetectPedestrianIntrusionAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectPedestrianIntrusionResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectPedestrianIntrusionReq := &DetectPedestrianIntrusionRequest{}
	openapiutil.Convert(request, detectPedestrianIntrusionReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectPedestrianIntrusionReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectPedestrianIntrusionResp, _err := client.DetectPedestrianIntrusionWithOptions(detectPedestrianIntrusionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectPedestrianIntrusionResp
	return _result, _err
}

func (client *Client) DetectVideoLivingFaceWithOptions(request *DetectVideoLivingFaceRequest, runtime *util.RuntimeOptions) (_result *DetectVideoLivingFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectVideoLivingFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectVideoLivingFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectVideoLivingFace(request *DetectVideoLivingFaceRequest) (_result *DetectVideoLivingFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectVideoLivingFaceResponse{}
	_body, _err := client.DetectVideoLivingFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectVideoLivingFaceAdvance(request *DetectVideoLivingFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectVideoLivingFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectVideoLivingFaceReq := &DetectVideoLivingFaceRequest{}
	openapiutil.Convert(request, detectVideoLivingFaceReq)
	if !tea.BoolValue(util.IsUnset(request.VideoUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectVideoLivingFaceReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectVideoLivingFaceResp, _err := client.DetectVideoLivingFaceWithOptions(detectVideoLivingFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectVideoLivingFaceResp
	return _result, _err
}

func (client *Client) EnhanceFaceWithOptions(request *EnhanceFaceRequest, runtime *util.RuntimeOptions) (_result *EnhanceFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnhanceFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnhanceFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnhanceFace(request *EnhanceFaceRequest) (_result *EnhanceFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnhanceFaceResponse{}
	_body, _err := client.EnhanceFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnhanceFaceAdvance(request *EnhanceFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *EnhanceFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	enhanceFaceReq := &EnhanceFaceRequest{}
	openapiutil.Convert(request, enhanceFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		enhanceFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	enhanceFaceResp, _err := client.EnhanceFaceWithOptions(enhanceFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = enhanceFaceResp
	return _result, _err
}

func (client *Client) ExtractFingerPrintWithOptions(request *ExtractFingerPrintRequest, runtime *util.RuntimeOptions) (_result *ExtractFingerPrintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageData)) {
		body["ImageData"] = request.ImageData
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtractFingerPrint"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtractFingerPrintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractFingerPrint(request *ExtractFingerPrintRequest) (_result *ExtractFingerPrintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractFingerPrintResponse{}
	_body, _err := client.ExtractFingerPrintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractFingerPrintAdvance(request *ExtractFingerPrintAdvanceRequest, runtime *util.RuntimeOptions) (_result *ExtractFingerPrintResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	extractFingerPrintReq := &ExtractFingerPrintRequest{}
	openapiutil.Convert(request, extractFingerPrintReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		extractFingerPrintReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	extractFingerPrintResp, _err := client.ExtractFingerPrintWithOptions(extractFingerPrintReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = extractFingerPrintResp
	return _result, _err
}

func (client *Client) ExtractPedestrianFeatureAttrWithOptions(request *ExtractPedestrianFeatureAttrRequest, runtime *util.RuntimeOptions) (_result *ExtractPedestrianFeatureAttrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		body["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtractPedestrianFeatureAttr"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtractPedestrianFeatureAttrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractPedestrianFeatureAttr(request *ExtractPedestrianFeatureAttrRequest) (_result *ExtractPedestrianFeatureAttrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtractPedestrianFeatureAttrResponse{}
	_body, _err := client.ExtractPedestrianFeatureAttrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractPedestrianFeatureAttrAdvance(request *ExtractPedestrianFeatureAttrAdvanceRequest, runtime *util.RuntimeOptions) (_result *ExtractPedestrianFeatureAttrResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	extractPedestrianFeatureAttrReq := &ExtractPedestrianFeatureAttrRequest{}
	openapiutil.Convert(request, extractPedestrianFeatureAttrReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		extractPedestrianFeatureAttrReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	extractPedestrianFeatureAttrResp, _err := client.ExtractPedestrianFeatureAttrWithOptions(extractPedestrianFeatureAttrReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = extractPedestrianFeatureAttrResp
	return _result, _err
}

func (client *Client) FaceBeautyWithOptions(request *FaceBeautyRequest, runtime *util.RuntimeOptions) (_result *FaceBeautyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Sharp)) {
		body["Sharp"] = request.Sharp
	}

	if !tea.BoolValue(util.IsUnset(request.Smooth)) {
		body["Smooth"] = request.Smooth
	}

	if !tea.BoolValue(util.IsUnset(request.White)) {
		body["White"] = request.White
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceBeauty"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceBeautyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceBeauty(request *FaceBeautyRequest) (_result *FaceBeautyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceBeautyResponse{}
	_body, _err := client.FaceBeautyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceBeautyAdvance(request *FaceBeautyAdvanceRequest, runtime *util.RuntimeOptions) (_result *FaceBeautyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	faceBeautyReq := &FaceBeautyRequest{}
	openapiutil.Convert(request, faceBeautyReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		faceBeautyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	faceBeautyResp, _err := client.FaceBeautyWithOptions(faceBeautyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = faceBeautyResp
	return _result, _err
}

func (client *Client) FaceFilterWithOptions(request *FaceFilterRequest, runtime *util.RuntimeOptions) (_result *FaceFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Strength)) {
		body["Strength"] = request.Strength
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceFilter"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceFilter(request *FaceFilterRequest) (_result *FaceFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceFilterResponse{}
	_body, _err := client.FaceFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceFilterAdvance(request *FaceFilterAdvanceRequest, runtime *util.RuntimeOptions) (_result *FaceFilterResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	faceFilterReq := &FaceFilterRequest{}
	openapiutil.Convert(request, faceFilterReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		faceFilterReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	faceFilterResp, _err := client.FaceFilterWithOptions(faceFilterReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = faceFilterResp
	return _result, _err
}

func (client *Client) FaceMakeupWithOptions(request *FaceMakeupRequest, runtime *util.RuntimeOptions) (_result *FaceMakeupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.MakeupType)) {
		body["MakeupType"] = request.MakeupType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Strength)) {
		body["Strength"] = request.Strength
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceMakeup"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceMakeupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceMakeup(request *FaceMakeupRequest) (_result *FaceMakeupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceMakeupResponse{}
	_body, _err := client.FaceMakeupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceMakeupAdvance(request *FaceMakeupAdvanceRequest, runtime *util.RuntimeOptions) (_result *FaceMakeupResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	faceMakeupReq := &FaceMakeupRequest{}
	openapiutil.Convert(request, faceMakeupReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		faceMakeupReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	faceMakeupResp, _err := client.FaceMakeupWithOptions(faceMakeupReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = faceMakeupResp
	return _result, _err
}

func (client *Client) FaceTidyupWithOptions(request *FaceTidyupRequest, runtime *util.RuntimeOptions) (_result *FaceTidyupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ShapeType)) {
		body["ShapeType"] = request.ShapeType
	}

	if !tea.BoolValue(util.IsUnset(request.Strength)) {
		body["Strength"] = request.Strength
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceTidyup"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceTidyupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceTidyup(request *FaceTidyupRequest) (_result *FaceTidyupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceTidyupResponse{}
	_body, _err := client.FaceTidyupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceTidyupAdvance(request *FaceTidyupAdvanceRequest, runtime *util.RuntimeOptions) (_result *FaceTidyupResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	faceTidyupReq := &FaceTidyupRequest{}
	openapiutil.Convert(request, faceTidyupReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		faceTidyupReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	faceTidyupResp, _err := client.FaceTidyupWithOptions(faceTidyupReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = faceTidyupResp
	return _result, _err
}

func (client *Client) GenRealPersonVerificationTokenWithOptions(request *GenRealPersonVerificationTokenRequest, runtime *util.RuntimeOptions) (_result *GenRealPersonVerificationTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateName)) {
		body["CertificateName"] = request.CertificateName
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateNumber)) {
		body["CertificateNumber"] = request.CertificateNumber
	}

	if !tea.BoolValue(util.IsUnset(request.MetaInfo)) {
		body["MetaInfo"] = request.MetaInfo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenRealPersonVerificationToken"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenRealPersonVerificationTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenRealPersonVerificationToken(request *GenRealPersonVerificationTokenRequest) (_result *GenRealPersonVerificationTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenRealPersonVerificationTokenResponse{}
	_body, _err := client.GenRealPersonVerificationTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateHumanAnimeStyleWithOptions(request *GenerateHumanAnimeStyleRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanAnimeStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgoType)) {
		query["AlgoType"] = request.AlgoType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateHumanAnimeStyle"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateHumanAnimeStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateHumanAnimeStyle(request *GenerateHumanAnimeStyleRequest) (_result *GenerateHumanAnimeStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateHumanAnimeStyleResponse{}
	_body, _err := client.GenerateHumanAnimeStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateHumanAnimeStyleAdvance(request *GenerateHumanAnimeStyleAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanAnimeStyleResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	generateHumanAnimeStyleReq := &GenerateHumanAnimeStyleRequest{}
	openapiutil.Convert(request, generateHumanAnimeStyleReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		generateHumanAnimeStyleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	generateHumanAnimeStyleResp, _err := client.GenerateHumanAnimeStyleWithOptions(generateHumanAnimeStyleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateHumanAnimeStyleResp
	return _result, _err
}

func (client *Client) GenerateHumanSketchStyleWithOptions(request *GenerateHumanSketchStyleRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanSketchStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnType)) {
		body["ReturnType"] = request.ReturnType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateHumanSketchStyle"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateHumanSketchStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateHumanSketchStyle(request *GenerateHumanSketchStyleRequest) (_result *GenerateHumanSketchStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateHumanSketchStyleResponse{}
	_body, _err := client.GenerateHumanSketchStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateHumanSketchStyleAdvance(request *GenerateHumanSketchStyleAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateHumanSketchStyleResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	generateHumanSketchStyleReq := &GenerateHumanSketchStyleRequest{}
	openapiutil.Convert(request, generateHumanSketchStyleReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		generateHumanSketchStyleReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	generateHumanSketchStyleResp, _err := client.GenerateHumanSketchStyleWithOptions(generateHumanSketchStyleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateHumanSketchStyleResp
	return _result, _err
}

func (client *Client) GetFaceEntityWithOptions(request *GetFaceEntityRequest, runtime *util.RuntimeOptions) (_result *GetFaceEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFaceEntity"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFaceEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFaceEntity(request *GetFaceEntityRequest) (_result *GetFaceEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFaceEntityResponse{}
	_body, _err := client.GetFaceEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRealPersonVerificationResultWithOptions(request *GetRealPersonVerificationResultRequest, runtime *util.RuntimeOptions) (_result *GetRealPersonVerificationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VerificationToken)) {
		body["VerificationToken"] = request.VerificationToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRealPersonVerificationResult"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRealPersonVerificationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRealPersonVerificationResult(request *GetRealPersonVerificationResultRequest) (_result *GetRealPersonVerificationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRealPersonVerificationResultResponse{}
	_body, _err := client.GetRealPersonVerificationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandPostureWithOptions(request *HandPostureRequest, runtime *util.RuntimeOptions) (_result *HandPostureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HandPosture"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HandPostureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandPosture(request *HandPostureRequest) (_result *HandPostureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandPostureResponse{}
	_body, _err := client.HandPostureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandPostureAdvance(request *HandPostureAdvanceRequest, runtime *util.RuntimeOptions) (_result *HandPostureResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	handPostureReq := &HandPostureRequest{}
	openapiutil.Convert(request, handPostureReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		handPostureReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	handPostureResp, _err := client.HandPostureWithOptions(handPostureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = handPostureResp
	return _result, _err
}

func (client *Client) LiquifyFaceWithOptions(request *LiquifyFaceRequest, runtime *util.RuntimeOptions) (_result *LiquifyFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.SlimDegree)) {
		body["SlimDegree"] = request.SlimDegree
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LiquifyFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LiquifyFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LiquifyFace(request *LiquifyFaceRequest) (_result *LiquifyFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LiquifyFaceResponse{}
	_body, _err := client.LiquifyFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LiquifyFaceAdvance(request *LiquifyFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *LiquifyFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	liquifyFaceReq := &LiquifyFaceRequest{}
	openapiutil.Convert(request, liquifyFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		liquifyFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	liquifyFaceResp, _err := client.LiquifyFaceWithOptions(liquifyFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = liquifyFaceResp
	return _result, _err
}

func (client *Client) ListFaceDbsWithOptions(request *ListFaceDbsRequest, runtime *util.RuntimeOptions) (_result *ListFaceDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["Offset"] = request.Offset
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFaceDbs"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFaceDbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaceDbs(request *ListFaceDbsRequest) (_result *ListFaceDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceDbsResponse{}
	_body, _err := client.ListFaceDbsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFaceEntitiesWithOptions(request *ListFaceEntitiesRequest, runtime *util.RuntimeOptions) (_result *ListFaceEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityIdPrefix)) {
		body["EntityIdPrefix"] = request.EntityIdPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFaceEntities"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFaceEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaceEntities(request *ListFaceEntitiesRequest) (_result *ListFaceEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFaceEntitiesResponse{}
	_body, _err := client.ListFaceEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeImageFaceWithOptions(request *MergeImageFaceRequest, runtime *util.RuntimeOptions) (_result *MergeImageFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MergeImageFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MergeImageFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MergeImageFace(request *MergeImageFaceRequest) (_result *MergeImageFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MergeImageFaceResponse{}
	_body, _err := client.MergeImageFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeImageFaceAdvance(request *MergeImageFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *MergeImageFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	mergeImageFaceReq := &MergeImageFaceRequest{}
	openapiutil.Convert(request, mergeImageFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		mergeImageFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	mergeImageFaceResp, _err := client.MergeImageFaceWithOptions(mergeImageFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = mergeImageFaceResp
	return _result, _err
}

func (client *Client) MonitorExaminationWithOptions(request *MonitorExaminationRequest, runtime *util.RuntimeOptions) (_result *MonitorExaminationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MonitorExamination"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MonitorExaminationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MonitorExamination(request *MonitorExaminationRequest) (_result *MonitorExaminationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MonitorExaminationResponse{}
	_body, _err := client.MonitorExaminationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MonitorExaminationAdvance(request *MonitorExaminationAdvanceRequest, runtime *util.RuntimeOptions) (_result *MonitorExaminationResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	monitorExaminationReq := &MonitorExaminationRequest{}
	openapiutil.Convert(request, monitorExaminationReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		monitorExaminationReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	monitorExaminationResp, _err := client.MonitorExaminationWithOptions(monitorExaminationReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = monitorExaminationResp
	return _result, _err
}

func (client *Client) PedestrianDetectAttributeWithOptions(request *PedestrianDetectAttributeRequest, runtime *util.RuntimeOptions) (_result *PedestrianDetectAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PedestrianDetectAttribute"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PedestrianDetectAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PedestrianDetectAttribute(request *PedestrianDetectAttributeRequest) (_result *PedestrianDetectAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PedestrianDetectAttributeResponse{}
	_body, _err := client.PedestrianDetectAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PedestrianDetectAttributeAdvance(request *PedestrianDetectAttributeAdvanceRequest, runtime *util.RuntimeOptions) (_result *PedestrianDetectAttributeResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	pedestrianDetectAttributeReq := &PedestrianDetectAttributeRequest{}
	openapiutil.Convert(request, pedestrianDetectAttributeReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		pedestrianDetectAttributeReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	pedestrianDetectAttributeResp, _err := client.PedestrianDetectAttributeWithOptions(pedestrianDetectAttributeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = pedestrianDetectAttributeResp
	return _result, _err
}

func (client *Client) QueryFaceImageTemplateWithOptions(request *QueryFaceImageTemplateRequest, runtime *util.RuntimeOptions) (_result *QueryFaceImageTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFaceImageTemplate"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFaceImageTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFaceImageTemplate(request *QueryFaceImageTemplateRequest) (_result *QueryFaceImageTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFaceImageTemplateResponse{}
	_body, _err := client.QueryFaceImageTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeActionWithOptions(request *RecognizeActionRequest, runtime *util.RuntimeOptions) (_result *RecognizeActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	if !tea.BoolValue(util.IsUnset(request.VideoData)) {
		body["VideoData"] = request.VideoData
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeAction"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeAction(request *RecognizeActionRequest) (_result *RecognizeActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeActionResponse{}
	_body, _err := client.RecognizeActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeActionAdvance(request *RecognizeActionAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeActionResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeActionReq := &RecognizeActionRequest{}
	openapiutil.Convert(request, recognizeActionReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := recognizeActionReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeActionReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeActionResp, _err := client.RecognizeActionWithOptions(recognizeActionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeActionResp
	return _result, _err
}

func (client *Client) RecognizeExpressionWithOptions(request *RecognizeExpressionRequest, runtime *util.RuntimeOptions) (_result *RecognizeExpressionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeExpression"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeExpressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeExpression(request *RecognizeExpressionRequest) (_result *RecognizeExpressionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeExpressionResponse{}
	_body, _err := client.RecognizeExpressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeExpressionAdvance(request *RecognizeExpressionAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeExpressionResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeExpressionReq := &RecognizeExpressionRequest{}
	openapiutil.Convert(request, recognizeExpressionReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeExpressionReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeExpressionResp, _err := client.RecognizeExpressionWithOptions(recognizeExpressionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeExpressionResp
	return _result, _err
}

func (client *Client) RecognizeFaceWithOptions(request *RecognizeFaceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Age)) {
		body["Age"] = request.Age
	}

	if !tea.BoolValue(util.IsUnset(request.Beauty)) {
		body["Beauty"] = request.Beauty
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		body["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Gender)) {
		body["Gender"] = request.Gender
	}

	if !tea.BoolValue(util.IsUnset(request.Glass)) {
		body["Glass"] = request.Glass
	}

	if !tea.BoolValue(util.IsUnset(request.Hat)) {
		body["Hat"] = request.Hat
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		body["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.MaxFaceNumber)) {
		body["MaxFaceNumber"] = request.MaxFaceNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Quality)) {
		body["Quality"] = request.Quality
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFace(request *RecognizeFaceRequest) (_result *RecognizeFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFaceResponse{}
	_body, _err := client.RecognizeFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFaceAdvance(request *RecognizeFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeFaceReq := &RecognizeFaceRequest{}
	openapiutil.Convert(request, recognizeFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeFaceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeFaceResp, _err := client.RecognizeFaceWithOptions(recognizeFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeFaceResp
	return _result, _err
}

func (client *Client) RecognizeHandGestureWithOptions(request *RecognizeHandGestureRequest, runtime *util.RuntimeOptions) (_result *RecognizeHandGestureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.GestureType)) {
		body["GestureType"] = request.GestureType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeHandGesture"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHandGestureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeHandGesture(request *RecognizeHandGestureRequest) (_result *RecognizeHandGestureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHandGestureResponse{}
	_body, _err := client.RecognizeHandGestureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeHandGestureAdvance(request *RecognizeHandGestureAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeHandGestureResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeHandGestureReq := &RecognizeHandGestureRequest{}
	openapiutil.Convert(request, recognizeHandGestureReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeHandGestureReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeHandGestureResp, _err := client.RecognizeHandGestureWithOptions(recognizeHandGestureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeHandGestureResp
	return _result, _err
}

func (client *Client) RecognizePublicFaceWithOptions(request *RecognizePublicFaceRequest, runtime *util.RuntimeOptions) (_result *RecognizePublicFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Task)) {
		body["Task"] = request.Task
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizePublicFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePublicFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizePublicFace(request *RecognizePublicFaceRequest) (_result *RecognizePublicFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePublicFaceResponse{}
	_body, _err := client.RecognizePublicFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizePublicFaceAdvance(request *RecognizePublicFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizePublicFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizePublicFaceReq := &RecognizePublicFaceRequest{}
	openapiutil.Convert(request, recognizePublicFaceReq)
	if !tea.BoolValue(util.IsUnset(request.Task)) {
		i0 := tea.Int(0)
		for _, item0 := range request.Task {
			if !tea.BoolValue(util.IsUnset(item0.ImageURLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.ImageURLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := recognizePublicFaceReq.Task[tea.IntValue(i0)]
				tmp.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	recognizePublicFaceResp, _err := client.RecognizePublicFaceWithOptions(recognizePublicFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizePublicFaceResp
	return _result, _err
}

func (client *Client) RetouchBodyWithOptions(request *RetouchBodyRequest, runtime *util.RuntimeOptions) (_result *RetouchBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.LengthenDegree)) {
		body["LengthenDegree"] = request.LengthenDegree
	}

	if !tea.BoolValue(util.IsUnset(request.SlimDegree)) {
		body["SlimDegree"] = request.SlimDegree
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RetouchBody"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetouchBodyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetouchBody(request *RetouchBodyRequest) (_result *RetouchBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetouchBodyResponse{}
	_body, _err := client.RetouchBodyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetouchBodyAdvance(request *RetouchBodyAdvanceRequest, runtime *util.RuntimeOptions) (_result *RetouchBodyResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	retouchBodyReq := &RetouchBodyRequest{}
	openapiutil.Convert(request, retouchBodyReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		retouchBodyReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	retouchBodyResp, _err := client.RetouchBodyWithOptions(retouchBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = retouchBodyResp
	return _result, _err
}

func (client *Client) RetouchSkinWithOptions(request *RetouchSkinRequest, runtime *util.RuntimeOptions) (_result *RetouchSkinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.RetouchDegree)) {
		body["RetouchDegree"] = request.RetouchDegree
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteningDegree)) {
		body["WhiteningDegree"] = request.WhiteningDegree
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RetouchSkin"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetouchSkinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetouchSkin(request *RetouchSkinRequest) (_result *RetouchSkinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetouchSkinResponse{}
	_body, _err := client.RetouchSkinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetouchSkinAdvance(request *RetouchSkinAdvanceRequest, runtime *util.RuntimeOptions) (_result *RetouchSkinResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	retouchSkinReq := &RetouchSkinRequest{}
	openapiutil.Convert(request, retouchSkinReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		retouchSkinReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	retouchSkinResp, _err := client.RetouchSkinWithOptions(retouchSkinReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = retouchSkinResp
	return _result, _err
}

func (client *Client) SearchFaceWithOptions(request *SearchFaceRequest, runtime *util.RuntimeOptions) (_result *SearchFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.DbNames)) {
		body["DbNames"] = request.DbNames
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.MaxFaceNum)) {
		body["MaxFaceNum"] = request.MaxFaceNum
	}

	if !tea.BoolValue(util.IsUnset(request.QualityScoreThreshold)) {
		body["QualityScoreThreshold"] = request.QualityScoreThreshold
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchFace"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFace(request *SearchFaceRequest) (_result *SearchFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchFaceResponse{}
	_body, _err := client.SearchFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFaceAdvance(request *SearchFaceAdvanceRequest, runtime *util.RuntimeOptions) (_result *SearchFaceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	searchFaceReq := &SearchFaceRequest{}
	openapiutil.Convert(request, searchFaceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		searchFaceReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	searchFaceResp, _err := client.SearchFaceWithOptions(searchFaceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = searchFaceResp
	return _result, _err
}

func (client *Client) SwapFacialFeaturesWithOptions(request *SwapFacialFeaturesRequest, runtime *util.RuntimeOptions) (_result *SwapFacialFeaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EditPart)) {
		body["EditPart"] = request.EditPart
	}

	if !tea.BoolValue(util.IsUnset(request.SourceImageData)) {
		body["SourceImageData"] = request.SourceImageData
	}

	if !tea.BoolValue(util.IsUnset(request.SourceImageURL)) {
		body["SourceImageURL"] = request.SourceImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageData)) {
		body["TargetImageData"] = request.TargetImageData
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageURL)) {
		body["TargetImageURL"] = request.TargetImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SwapFacialFeatures"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwapFacialFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwapFacialFeatures(request *SwapFacialFeaturesRequest) (_result *SwapFacialFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwapFacialFeaturesResponse{}
	_body, _err := client.SwapFacialFeaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwapFacialFeaturesAdvance(request *SwapFacialFeaturesAdvanceRequest, runtime *util.RuntimeOptions) (_result *SwapFacialFeaturesResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	swapFacialFeaturesReq := &SwapFacialFeaturesRequest{}
	openapiutil.Convert(request, swapFacialFeaturesReq)
	if !tea.BoolValue(util.IsUnset(request.SourceImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.SourceImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		swapFacialFeaturesReq.SourceImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.TargetImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.TargetImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		swapFacialFeaturesReq.TargetImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	swapFacialFeaturesResp, _err := client.SwapFacialFeaturesWithOptions(swapFacialFeaturesReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = swapFacialFeaturesResp
	return _result, _err
}

func (client *Client) UpdateFaceEntityWithOptions(request *UpdateFaceEntityRequest, runtime *util.RuntimeOptions) (_result *UpdateFaceEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFaceEntity"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFaceEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFaceEntity(request *UpdateFaceEntityRequest) (_result *UpdateFaceEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFaceEntityResponse{}
	_body, _err := client.UpdateFaceEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyFaceMaskWithOptions(request *VerifyFaceMaskRequest, runtime *util.RuntimeOptions) (_result *VerifyFaceMaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageData)) {
		body["ImageData"] = request.ImageData
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.RefData)) {
		body["RefData"] = request.RefData
	}

	if !tea.BoolValue(util.IsUnset(request.RefUrl)) {
		body["RefUrl"] = request.RefUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyFaceMask"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyFaceMaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyFaceMask(request *VerifyFaceMaskRequest) (_result *VerifyFaceMaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyFaceMaskResponse{}
	_body, _err := client.VerifyFaceMaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyFaceMaskAdvance(request *VerifyFaceMaskAdvanceRequest, runtime *util.RuntimeOptions) (_result *VerifyFaceMaskResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("facebody"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	verifyFaceMaskReq := &VerifyFaceMaskRequest{}
	openapiutil.Convert(request, verifyFaceMaskReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		verifyFaceMaskReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.RefUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.RefUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		verifyFaceMaskReq.RefUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	verifyFaceMaskResp, _err := client.VerifyFaceMaskWithOptions(verifyFaceMaskReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = verifyFaceMaskResp
	return _result, _err
}
